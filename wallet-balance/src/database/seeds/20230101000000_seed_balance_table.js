module.exports = {
   up: async (queryInterface, Sequelize) => {
     await queryInterface.bulkInsert('balances', [
       { accountId: 1, balance: 100.50 },
       { accountId: 2, balance: 250.00 },
       { accountId: 3, balance: 999.99 }
     ]);
   },
   down: async (queryInterface, Sequelize) => {
     await queryInterface.bulkDelete('balances', null, {});
   }
 };
 