const path = require('path');
const { Sequelize } = require('sequelize');
const { sequelize } = require('../models/balance.model');
const { Umzug, SequelizeStorage } = require('umzug');

const seeder = new Umzug({
  migrations: {
    glob: path.join(__dirname, 'seeds/*.js'),
    resolve: ({ name, path, context }) => {
      const seed = require(path);
      return {
        name,
        up: async () => seed.up(context, Sequelize),
        down: async () => seed.down(context, Sequelize),
      };
    },
  },
  context: sequelize.getQueryInterface(),
  storage: new SequelizeStorage({
    sequelize,
    tableName: 'SequelizeSeedMeta',
  }),
  logger: console,
});

async function runSeeds() {
  try {
    const seeds = await seeder.up();
    console.log('Seeds executed:', seeds.map((s) => s.name));
  } catch (error) {
    console.error('Error executing seeds:', error);
    process.exit(1);
  }
}

if (require.main === module) {
  runSeeds();
}

module.exports = { runSeeds };
