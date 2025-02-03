require('dotenv').config();
const { Kafka } = require('kafkajs');
const { Balance } = require('./models/balance.model');

const kafka = new Kafka({
  clientId: process.env.KAFKA_CLIENT_ID,
  brokers: [process.env.KAFKA_BROKER]
});

const consumer = kafka.consumer({ groupId: process.env.KAFKA_GROUP_ID });

exports.startKafkaConsumer = async () => {
  try {
    await consumer.connect();
    await consumer.subscribe({ topic: process.env.KAFKA_BALANCE_UPDATES_TOPIC, fromBeginning: true });

    console.log('afka consumer connected and subscribed to the topic:', process.env.KAFKA_BALANCE_UPDATES_TOPIC);

    await consumer.run({
      eachMessage: async ({ topic, partition, message }) => {
        try {
          const eventData = JSON.parse(message.value.toString());
          console.log(`Message received from topic ${topic}`, eventData);
    
          const { Name, Payload } = eventData;
          
          if (Name === "BalanceUpdated") {
            const {
              account_id_from,
              account_id_to,
              balance_account_id_from,
              balance_account_id_to
            } = Payload;
    
            await Balance.upsert({
              accountId: account_id_from,
              balance: balance_account_id_from
            });
    
            await Balance.upsert({
              accountId: account_id_to,
              balance: balance_account_id_to
            });
    
            console.log(`Balances successfully updated:
              Account: ${account_id_from} => ${balance_account_id_from}
              Account: ${account_id_to} => ${balance_account_id_to}
            `);
          } else {
            console.log("Unrecognized or unhandled event:", Name);
          }
        } catch (err) {
          console.error('Error processing Kafka message:', err);
        }
      }
    });
    
  } catch (error) {
    console.error('Error on Kafka consumer:', error);
  }
};
