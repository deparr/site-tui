import { RGBA, SyntaxStyle } from "@opentui/core";

export const colors: Record<"light" | "dark", Record<string, RGBA>> = {
    light: {},
    dark: {
        background: RGBA.fromHex("#151515"),
        text: RGBA.fromHex("#a5a8a6"),
        border: RGBA.fromHex("#696969"),
        accent: RGBA.fromHex("#e78c45"),
        highlight: RGBA.fromHex("#e5e8e6"),
        dim: RGBA.fromHex("#696969"),
        error: RGBA.fromHex("#cc6666"),
    },

};

export const syntaxColors: Record<"light" | "dark", Record<string, RGBA>> = {
    dark: {
        white: RGBA.fromHex("#d8d8d8"),
        off_white: RGBA.fromHex("#a8a8a8"),
        gray: RGBA.fromHex("#7b7b7b"),
        red: RGBA.fromHex("#ac4242"),
        green: RGBA.fromHex("#90a959"),
        yellow: RGBA.fromHex("#f4bf75"),
        blue: RGBA.fromHex("#6a9fb5"),
        purple: RGBA.fromHex("#aa759f"),
        cyan: RGBA.fromHex("#75b5aa"),
        orange: RGBA.fromHex("#cc7f40"),
        bg: RGBA.fromHex("#262626"),
    },
    light: {
        black: RGBA.fromHex("#393842"),
        gray: RGBA.fromHex("#9fa0a6"),
        purple: RGBA.fromHex("#a625a4"),
        blue: RGBA.fromHex("#4078f1"),
        cyan: RGBA.fromHex("#0083bb"),
        green: RGBA.fromHex("#50a04f"),
        yellow: RGBA.fromHex("#b66a00"),
        orange: RGBA.fromHex("#e35549"),
        red: RGBA.fromHex("#c91142"),
        bg: RGBA.fromHex("#f4f4f4"),
    }
};

export const tui = {
    syntax: {
        dark: SyntaxStyle.fromStyles({
            attribute: {
                fg: syntaxColors.dark.blue,
            },
            "attribute.builtin": {
                fg: syntaxColors.dark.purple,
            },
            boolean: {
                fg: syntaxColors.dark.orange,
            },
            character: {
                fg: syntaxColors.dark.green,
            },
            "character.special": {
                fg: syntaxColors.dark.white,
            },
            comment: {
                fg: syntaxColors.dark.gray,
            },
            "comment.documentation": {
                fg: syntaxColors.dark.gray,
            },
            "comment.error": {
                fg: syntaxColors.dark.red,
            },
            "comment.note": {
                fg: syntaxColors.dark.gray,
            },
            "comment.todo": {
                fg: syntaxColors.dark.gray,
            },
            "comment.warning": {
                fg: syntaxColors.dark.gray,
            },
            constant: {
                fg: syntaxColors.dark.orange,
            },
            "constant.builtin": {
                fg: syntaxColors.dark.orange,
            },
            "constant.macro": {
                fg: syntaxColors.dark.yellow,
            },
            constructor: {
                fg: syntaxColors.dark.white,
            },
            "diff.delta": {
                fg: syntaxColors.dark.blue,
            },
            "diff.minus": {
                fg: syntaxColors.dark.red,
            },
            "diff.plus": {
                fg: syntaxColors.dark.green,
            },
            function: {
                fg: syntaxColors.dark.blue,
            },
            "function.builtin": {
                fg: syntaxColors.dark.cyan,
            },
            "function.call": {
                fg: syntaxColors.dark.white,
            },
            "function.macro": {
                fg: syntaxColors.dark.cyan,
            },
            "function.method": {
                fg: syntaxColors.dark.blue,
            },
            "function.method.call": {
                fg: syntaxColors.dark.white,
            },
            keyword: {
                fg: syntaxColors.dark.purple,
            },
            "keyword.conditional": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.conditional.ternary": {
                fg: syntaxColors.dark.cyan,
            },
            "keyword.coroutine": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.debug": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.directive": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.directive.define": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.exception": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.function": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.import": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.modifier": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.operator": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.repeat": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.return": {
                fg: syntaxColors.dark.purple,
            },
            "keyword.type": {
                fg: syntaxColors.dark.purple,
            },
            label: {
                fg: syntaxColors.dark.yellow,
            },
            "markup.heading": {
                bold: true,
                fg: syntaxColors.dark.white,
            },
            "markup.heading.1": {
                bold: true,
                fg: syntaxColors.dark.white,
            },
            "markup.heading.2": {
                bold: true,
                fg: syntaxColors.dark.white,
            },
            "markup.heading.3": {
                bold: true,
                fg: syntaxColors.dark.white,
            },
            "markup.heading.4": {
                bold: true,
                fg: syntaxColors.dark.white,
            },
            "markup.heading.5": {
                bold: true,
                fg: syntaxColors.dark.white,
            },
            "markup.heading.6": {
                bold: true,
                fg: syntaxColors.dark.white,
            },
            "markup.italic": {
                italic: true
            },
            "markup.link": {
                fg: syntaxColors.dark.blue,
            },
            "markup.link.label": {
                fg: syntaxColors.dark.orange,
            },
            "markup.link.url": {
                underline: true
            },
            "markup.list": {
                fg: syntaxColors.dark.white,
            },
            "markup.list.checked": {
                fg: syntaxColors.dark.green,
            },
            "markup.list.unchecked": {
                fg: syntaxColors.dark.yellow,
            },
            "markup.math": {
                fg: syntaxColors.dark.cyan,
            },
            "markup.quote": {
                fg: syntaxColors.dark.gray,
            },
            "markup.raw": {
                fg: syntaxColors.dark.white,
            },
            "markup.raw.block": {
                fg: syntaxColors.dark.white,
            },
            // "markup.strikethrough": {
            //     strikethrough: true
            // },
            "markup.strong": {
                bold: true
            },
            "markup.underline": {
                underline: true
            },
            module: {
                fg: syntaxColors.dark.white,
            },
            "module.builtin": {
                fg: syntaxColors.dark.white,
            },
            number: {
                fg: syntaxColors.dark.orange,
            },
            "number.float": {
                fg: syntaxColors.dark.orange,
            },
            operator: {
                fg: syntaxColors.dark.off_white,
            },
            property: {
                fg: syntaxColors.dark.white,
            },
            "punctuation.bracket": {
                fg: syntaxColors.dark.off_white,
            },
            "punctuation.delimiter": {
                fg: syntaxColors.dark.off_white,
            },
            "punctuation.special": {
                fg: syntaxColors.dark.orange,
            },
            string: {
                fg: syntaxColors.dark.green,
            },
            "string.documentation": {
                fg: syntaxColors.dark.green,
            },
            "string.escape": {
                fg: syntaxColors.dark.red,
            },
            "string.regexp": {
                fg: syntaxColors.dark.red,
            },
            "string.special": {
                fg: syntaxColors.dark.red,
            },
            "string.special.path": {
                fg: syntaxColors.dark.red,
            },
            "string.special.symbol": {
                fg: syntaxColors.dark.red,
            },
            "string.special.url": {
                fg: syntaxColors.dark.cyan,
            },
            tag: {
                fg: syntaxColors.dark.blue,
            },
            "tag.attribute": {
                fg: syntaxColors.dark.yellow,
            },
            "tag.builtin": {
                fg: syntaxColors.dark.blue,
            },
            "tag.delimiter": {
                fg: syntaxColors.dark.off_white,
            },
            type: {
                fg: syntaxColors.dark.white,
            },
            "type.builtin": {
                fg: syntaxColors.dark.yellow,
            },
            "type.definition": {
                bold: false,
                fg: syntaxColors.dark.cyan,
            },
            variable: {
                fg: syntaxColors.dark.white,
            },
            "variable.builtin": {
                fg: syntaxColors.dark.white,
            },
            "variable.member": {
                fg: syntaxColors.dark.white,
            },
            "variable.parameter": {
                fg: syntaxColors.dark.white,
            },
            "variable.parameter.builtin": {
                fg: syntaxColors.dark.white,
            }
        }),
        light: SyntaxStyle.fromStyles({
            attribute: {
                fg: syntaxColors.light.purple,
            },
            "attribute.builtin": {
                fg: syntaxColors.light.purple,
            },
            boolean: {
                fg: syntaxColors.light.yellow,
            },
            character: {
                fg: syntaxColors.light.green,
            },
            "character.special": {
                fg: syntaxColors.light.black,
            },
            comment: {
                fg: syntaxColors.light.gray,
            },
            "comment.documentation": {
                fg: syntaxColors.light.gray,
            },
            "comment.error": {
                fg: syntaxColors.light.gray,
            },
            "comment.note": {
                fg: syntaxColors.light.gray,
            },
            "comment.todo": {
                fg: syntaxColors.light.gray,
            },
            "comment.warning": {
                fg: syntaxColors.light.gray,
            },
            constant: {
                fg: syntaxColors.light.yellow,
            },
            "constant.builtin": {
                fg: syntaxColors.light.orange,
            },
            "constant.macro": {
                fg: syntaxColors.light.red,
            },
            constructor: {
                fg: syntaxColors.light.black,
            },
            "diff.delta": {
                fg: syntaxColors.light.cyan,
            },
            "diff.minus": {
                fg: syntaxColors.light.red,
            },
            "diff.plus": {
                fg: syntaxColors.light.green,
            },
            function: {
                fg: syntaxColors.light.blue,
            },
            "function.builtin": {
                fg: syntaxColors.light.cyan,
            },
            "function.call": {
                fg: syntaxColors.light.black,
            },
            "function.macro": {
                fg: syntaxColors.light.red,
            },
            "function.method": {
                fg: syntaxColors.light.blue,
            },
            "function.method.call": {
                fg: syntaxColors.light.black,
            },
            keyword: {
                fg: syntaxColors.light.purple,
            },
            "keyword.conditional": {
                fg: syntaxColors.light.purple,
            },
            "keyword.conditional.ternary": {
                fg: syntaxColors.light.black,
            },
            "keyword.coroutine": {
                fg: syntaxColors.light.purple,
            },
            "keyword.debug": {
                fg: syntaxColors.light.purple,
            },
            "keyword.directive": {
                fg: syntaxColors.light.purple,
            },
            "keyword.directive.define": {
                fg: syntaxColors.light.purple,
            },
            "keyword.exception": {
                fg: syntaxColors.light.purple,
            },
            "keyword.function": {
                fg: syntaxColors.light.purple,
            },
            "keyword.import": {
                fg: syntaxColors.light.purple,
            },
            "keyword.modifier": {
                fg: syntaxColors.light.purple,
            },
            "keyword.operator": {
                fg: syntaxColors.light.purple,
            },
            "keyword.repeat": {
                fg: syntaxColors.light.purple,
            },
            "keyword.return": {
                fg: syntaxColors.light.purple,
            },
            "keyword.type": {
                fg: syntaxColors.light.purple,
            },
            label: {
                bold: true,
                fg: syntaxColors.light.black,
            },
            "markup.heading": {
                bold: true,
                fg: syntaxColors.light.black,
            },
            "markup.heading.1": {
                bold: true,
                fg: syntaxColors.light.black,
            },
            "markup.heading.2": {
                bold: true,
                fg: syntaxColors.light.black,
            },
            "markup.heading.3": {
                bold: true,
                fg: syntaxColors.light.black,
            },
            "markup.heading.4": {
                bold: true,
                fg: syntaxColors.light.black,
            },
            "markup.heading.5": {
                bold: true,
                fg: syntaxColors.light.black,
            },
            "markup.heading.6": {
                bold: true,
                fg: syntaxColors.light.black,
            },
            "markup.italic": {
                italic: true
            },
            "markup.link": {
                fg: syntaxColors.light.green,
            },
            "markup.link.label": {
                fg: syntaxColors.light.black,
            },
            "markup.link.url": {
                underline: true
            },
            "markup.list": {
                fg: syntaxColors.light.black,
            },
            "markup.list.checked": {
                fg: syntaxColors.light.green,
            },
            "markup.list.unchecked": {
                fg: syntaxColors.light.gray,
            },
            "markup.math": {
                bold: true,
                fg: syntaxColors.light.blue,
            },
            "markup.quote": {
                fg: syntaxColors.light.green,
            },
            "markup.raw": {
                fg: syntaxColors.light.black,
            },
            "markup.raw.block": {
                fg: syntaxColors.light.black,
            },
            // "markup.strikethrough": {
            //     strikethrough: true
            // },
            "markup.strong": {
                bold: true
            },
            "markup.underline": {
                underline: true
            },
            module: {
                fg: syntaxColors.light.black,
            },
            "module.builtin": {
                fg: syntaxColors.light.cyan,
            },
            number: {
                fg: syntaxColors.light.yellow,
            },
            "number.float": {
                fg: syntaxColors.light.yellow,
            },
            operator: {
                fg: syntaxColors.light.blue,
            },
            property: {
                fg: syntaxColors.light.black,
            },
            "punctuation.bracket": {
                fg: syntaxColors.light.black,
            },
            "punctuation.delimiter": {
                fg: syntaxColors.light.black,
            },
            "punctuation.special": {
                fg: syntaxColors.light.yellow,
            },
            string: {
                fg: syntaxColors.light.green,
            },
            "string.documentation": {
                fg: syntaxColors.light.green,
            },
            "string.escape": {
                fg: syntaxColors.light.red,
            },
            "string.regexp": {
                fg: syntaxColors.light.green,
            },
            "string.special": {
                fg: syntaxColors.light.green,
            },
            "string.special.path": {
                fg: syntaxColors.light.green,
            },
            "string.special.symbol": {
                fg: syntaxColors.light.green,
            },
            "string.special.url": {
                fg: syntaxColors.light.cyan,
            },
            tag: {
                fg: syntaxColors.light.blue,
            },
            "tag.attribute": {
                fg: syntaxColors.light.yellow,
            },
            "tag.builtin": {
                fg: syntaxColors.light.blue,
            },
            "tag.delimiter": {
                fg: syntaxColors.light.black,
            },
            type: {
                fg: syntaxColors.light.yellow,
            },
            "type.builtin": {
                fg: syntaxColors.light.yellow,
            },
            "type.definition": {
                fg: syntaxColors.light.yellow,
            },
            variable: {
                fg: syntaxColors.light.black,
            },
            "variable.builtin": {
                fg: syntaxColors.light.cyan,
            },
            "variable.member": {
                fg: syntaxColors.light.black,
            },
            "variable.parameter": {
                fg: syntaxColors.light.black,
            },
            "variable.parameter.builtin": {
                fg: syntaxColors.light.cyan,
            }
        })
    }
};
