const { runMigrations } = require('./migrate');
const { runSeeds } = require('./seed');

async function runMigrationsAndSeeds() {
  await runMigrations();
  await runSeeds();
}

module.exports = { runMigrationsAndSeeds };
