const path = require('path');
const { Sequelize } = require('sequelize');
const { sequelize } = require('../models/balance.model');
const { Umzug, SequelizeStorage } = require('umzug');

const migrator = new Umzug({
  migrations: {
    glob: path.join(__dirname, 'migrations/*.js'),

    resolve: ({ name, path, context }) => {
      const migration = require(path);
      return {
        name,
        up: async () => migration.up(context, Sequelize),
        down: async () => migration.down(context, Sequelize),
      };
    },
  },

  context: sequelize.getQueryInterface(),

  storage: new SequelizeStorage({
    sequelize,
    tableName: 'SequelizeMeta',
  }),

  logger: console, 
});

async function runMigrations() {
  try {
    const migrations = await migrator.up();
    console.log(
      'Migrations executed:',
      migrations.map((m) => m.name)
    );
  } catch (error) {
    console.error('Error executing migrations:', error);
    process.exit(1);
  }
}

if (require.main === module) {
  runMigrations();
}

module.exports = { runMigrations };
