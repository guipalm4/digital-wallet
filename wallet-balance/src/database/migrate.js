const path = require('path');
const { Sequelize } = require('sequelize');
const { sequelize } = require('../models/balance.model');
const { Umzug, SequelizeStorage } = require('umzug');

const migrator = new Umzug({
  migrations: {
    // Usamos "glob" para indicar onde estão nossas migrations
    glob: path.join(__dirname, 'migrations/*.js'),

    // "resolve" customizado para passar (queryInterface, Sequelize) nos métodos up/down
    resolve: ({ name, path, context }) => {
      const migration = require(path);
      return {
        name,
        up: async () => migration.up(context, Sequelize),
        down: async () => migration.down(context, Sequelize),
      };
    },
  },
  // O "context" passado para up/down será o queryInterface do Sequelize
  context: sequelize.getQueryInterface(),

  // Define onde o Umzug salva o status das migrations (por padrão, tabela SequelizeMeta)
  storage: new SequelizeStorage({
    sequelize,
    tableName: 'SequelizeMeta', // Padrão ou outro nome de sua preferência
  }),

  logger: console, // Opcional: mostra logs no console
});

async function runMigrations() {
  try {
    const migrations = await migrator.up();
    console.log(
      'Migrations executadas:',
      migrations.map((m) => m.name)
    );
  } catch (error) {
    console.error('Erro ao executar migrations:', error);
    process.exit(1);
  }
}

// Se chamar este arquivo diretamente (e.g. "node migrate.js"), executa as migrations
if (require.main === module) {
  runMigrations();
}

module.exports = { runMigrations };
