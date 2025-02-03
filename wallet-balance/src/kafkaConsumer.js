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

    console.log('Consumidor Kafka conectado e subscrito ao tópico:', process.env.KAFKA_BALANCE_UPDATES_TOPIC);

    await consumer.run({
      eachMessage: async ({ topic, partition, message }) => {
        try {
          const eventData = JSON.parse(message.value.toString());
          console.log(`Mensagem recebida do tópico ${topic}`, eventData);

          const { accountId, balance } = eventData;

          if (!accountId || balance === undefined) {
            console.log('Mensagem inválida, ignorando...');
            return;
          }

          await Balance.upsert({
            accountId: accountId,
            balance: balance
          });

          console.log(`Balance da conta ${accountId} atualizado para: ${balance}`);
        } catch (err) {
          console.error('Erro ao processar mensagem Kafka:', err);
        }
      }
    });
  } catch (error) {
    console.error('Erro no Kafka consumer:', error);
  }
};
