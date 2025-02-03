require('dotenv').config();
const express = require('express');
const routes = require('./routes');
const { runMigrationsAndSeeds } = require('./database/setup');
const { startKafkaConsumer } = require('./kafkaConsumer');

const app = express();
const PORT = process.env.PORT || 3003;

app.use(express.json());
app.use(routes);

async function init() {
  try {
    await runMigrationsAndSeeds();

    startKafkaConsumer();

    app.listen(PORT, () => {
      console.log(`Wallet Balances service running on port ${PORT}`);
    });
  } catch (error) {
    console.error('Error starting the balance service', error);
    process.exit(1);
  }
}

init();
