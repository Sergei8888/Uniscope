module.exports = {
    meta: {
        type: 'problem',
        docs: {
            description: 'enforce Page suffix in files inside pages directory',
            category: 'Stylistic Issues',
            recommended: false,
        },
        fixable: null,
        schema: [],
    },
    create(context) {
        const filename = context.getFilename();
        const path = require('path');
        if (filename.includes(path.sep + 'pages' + path.sep)) {
            if (!filename.endsWith('Page' + path.extname(filename))) {
                context.report({
                    node: null,
                    message: `Files inside the pages directory should have a Page suffix.`,
                });
            }
        }
        return {};
    },
};
