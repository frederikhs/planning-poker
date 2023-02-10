module.exports = {
    content: [
        "./src/**/*.{js,jsx,ts,tsx}",
    ],
    corePlugins: {
        container: false
    },
    plugins: [
        function ({addComponents}) {
            addComponents({
                '.container': {
                    maxWidth: '100%',
                    '@screen sm': {
                        maxWidth: '640px',
                    },
                    '@screen md': {
                        maxWidth: '768px',
                    },
                    '@screen lg': {
                        maxWidth: '1280px',
                    },
                    '@screen xl': {
                        maxWidth: '1280px',
                    },
                }
            })
        }
    ]
}