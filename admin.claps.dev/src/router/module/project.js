const projectRouter = [
    {
        path: '/project',
        name: 'project',
        meta: {
            auth: true,
        },
        component: () => import('../../views/project/Project.vue')
    },
    {
        path: '/project/:name',
        name: 'project/detail',
        meta: {
            auth: true,
        },
        component: () => import('../../views/project/Detail.vue')
    },
    {
        path: '/project/:name/addMember',
        name: 'addMember',
        meta: {
            auth: true,
        },
        component: () => import('../../views/project/AddMember.vue')
    },
    {
        path: '/project/:name/addRepository',
        name: 'addRepository',
        meta: {
            auth: true,
        },
        component: () => import('../../views/project/AddRepository.vue')
    },
    {
        path: '/project/:name/transact',
        name: 'projectTransaction',
        meta: {
            auth: true,
        },
        component: () => import('../../views/project/ProjectTransaction.vue')
    },
    {
        path: '/projects/addProject',
        name: 'addProject',
        meta: {
            auth: true,
        },
        component: () => import('../../views/project/AddProject.vue')
    },
];

export default projectRouter;
