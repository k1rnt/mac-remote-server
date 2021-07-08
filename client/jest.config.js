module.exports = {
    transform: {
        '.*\\.(ts)$': '<rootDir>/node_modules/ts-jest',    // TypeScriptファイルをテストする場合
    },
    moduleFileExtensions: ['ts', 'js'], // テスト対象の拡張子を列挙する
    moduleNameMapper: {
        "^~/(.+)": "<rootDir>/$1",
        "^@/(.+)": "<rootDir>/$1"
    }
};