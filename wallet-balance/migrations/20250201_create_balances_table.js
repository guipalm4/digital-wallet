exports.up = function(knex) {
   return knex.schema.createTable('balances', function(table) {
     table.string('account_id').primary();
     table.float('balance').notNullable();
   });
 };
 
 exports.down = function(knex) {
   return knex.schema.dropTable('balances');
 };
 