module.exports = {
    "env": {
        "browser": true,
        "es2021": true
    },
    "extends": [],
    "overrides": [
        {
            "env": {
                "node": true
            },
            "files": [
                ".eslintrc.{js,cjs}"
            ],
            "parserOptions": {
                "sourceType": "script"
            }
        }
    ],
    "parser": "@typescript-eslint/parser",
    "parserOptions": {
        "ecmaVersion": "latest",
        "sourceType": "module"
    },
    "plugins": [
        "@typescript-eslint",
        "react"
    ],
    "rules": {

        "no-dupe-else-if": "error",
        "no-use-before-define": "error",
        "use-isnan": "warn",
        "no-dupe-keys": "error",
        "no-var": "error",

    }
}
