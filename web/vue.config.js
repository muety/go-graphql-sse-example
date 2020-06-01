module.exports = {
    outputDir: './dist',
    devServer: {
        port: 4000,
        proxy: {
            '^[/api|/static]': {
                target: 'http://localhost:8080'
            }
        }
    },
    lintOnSave: true
}
