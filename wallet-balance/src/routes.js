const { Router } = require('express');
const { getBalance } = require('./controllers/balance.controller');

const router = Router();

router.get('/balances/:account_id', getBalance);

module.exports = router;
