import mock from 'mockjs'

mock.setup({timeout:'20-200'})

mock.mock(/auth\/login/, {
    code: 2000,
    msg: '',
    data: 'mock token',
})

mock.mock(/auth\/info/, {
    code: 2000,
    msg: '',
    data: {
        id: 1,
        acct: 'mock_acct',
        account_status: 'NORMAL',
        user_type: 'admin',
        alias: 'mock user'
    },
})

mock.mock(/user\/list/, {
    code: 2000,
    msg: '',
    data: {
        count: 10,
        list: [{
            id: 1,
            acct: 'mock_acct',
            account_status: 'NORMAL',
            user_type: 'admin',
            alias: 'mock user',
            create_at: 1665561443
        }]
    },
})