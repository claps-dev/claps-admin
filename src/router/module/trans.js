const transRouter = [
    {
        path: '/trans',
        name: 'trans',
        meta: {
            auth: true,
        },
        component: () => import('../../views/trans/Trans.vue'),
    },

    {
        path: '/transactions',
        name: "transactions",
        meta: {
            auth: true,
        },
        component: () => import('../../views/trans/Transactions.vue'),
    },
    {
        path: '/transfers',
        name: "transfers",
        meta: {
            auth: true,
        },
        component: () => import('../../views/trans/Transfers.vue'),
    },
];

export default transRouter;
