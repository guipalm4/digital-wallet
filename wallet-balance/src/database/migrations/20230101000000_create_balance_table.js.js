module.exports = {
  up: async (queryInterface, Sequelize) => {
    await queryInterface.createTable('balances', {
      id: {
        type: Sequelize.INTEGER,
        autoIncrement: true,
        primaryKey: true
      },
      accountId: {
        type: Sequelize.STRING,
        allowNull: false,
        unique: true
      },
      balance: {
        type: Sequelize.DECIMAL(12, 2),
        allowNull: false,
        defaultValue: 0.00
      }
    });
  },

  down: async (queryInterface, Sequelize) => {
    await queryInterface.dropTable('balances');
  }
};
