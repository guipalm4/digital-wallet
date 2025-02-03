const { Balance } = require('../models/balance.model');


exports.getBalance = async (req, res) => {
  try {
    const { account_id } = req.params;
    const balanceRecord = await Balance.findOne({ where: { accountId: account_id } });

    if (!balanceRecord) {
      return res.status(404).json({ error: 'Account not found.' });
    }

    return res.json({
      accountId: balanceRecord.accountId,
      balance: balanceRecord.balance
    });
  } catch (error) {
    console.error('Failed to retrieve balance:', error);
    return res.status(500).json({ error: 'Internal Server Error' });
  }
};
