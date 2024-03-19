# Introduction

Fantasy Name Generator has two interfaces that each strike its own balance between ease of use and functionality. The Simple Interface is extremely easy but only offers you a limited form of control over the names that get generated. The Advanced Interface is not quite as easy to learn -- see instructions for using it below -- but allows you practically limitless control over the type of names that get generated. Among other things, it allows you to generate names that are variations on a name you already have in mind.

# Simple Interface

The Simple Interface provides a drop-down box to allow you to set what kind of names you want to generate. The "Default Names" choice in this list contains a diverse mix of generic name types, while the other choices are more specialized.

To use the Simple Interface, simply select a name type from the drop-down list and hit the Generate Names button next to it.

# Advanced Interface

The Advanced Interface offers a maximum of control over the names that get generated. You can use it to generate names or words with unfamiliar linguistic roots or to generate variations on a specific name.

The Advanced Interface has a "Collapse Triples" checkbox, which is checked by default. When on, names with three or more of the same letter in a row will be reduced reduced to just two prior to being printed out. So if the name "Purghinnn" was generated and the "Collapse Triples" checkbox was checked, it would appear as just "Purghinn." In addition, if the letter in question does not make sense to have even doubled, it will be reduced to just one. So the name "Leguiin" would be reduced to just "Leguin." If the "Collapse Triples" checkbox is left unchecked, this reduction step will not take place. It doesn't mean you will get unnatural double and triple letter sequences; it just means no efforts will be made to prevent them. We recommend you leave this option turned on unless you have a special need for it to be off.

The heart of the Advanced Interface is the "Name Generation Template" field. The name generation template describes how names should be constructed. For instance, a template could dictate that all names be a consonant followed by a vowel or vowel combination followed by a consonant or consonant combination. This template would generate short names. The template could also employ the use of random predefined syllables (such as 'tor', 'ash', 'ald', and 'dar'); a very effective template, one used frequently when default names are generated using the Simple Interface, is one that dictates names should consist simply of two predefined syllables.

Templates get more complicated than that, but they don't have to. Simple and effective templates may do nothing but list how many consonant combinations, vowel combinations, and syllables should comprise a word, and in what order. At generation time, random consonant combinations, random vowel combinations, and random syllables will be placed together according to the template to form each name. The randomizer used will be weighted so that more frequently seen letters, such as 't' and 's', will appear more often than less frequently seen letters, such as 'j' and 'z'. This goes a long way toward generating good, pronouncable names.

A consonant combination is two or three consonants that tend to go well together, such as 'ch', 'ck', 'tt', and 'str'. Examples of vowel combinations are 'ou', 'ai', and 'ee'. In your template, you will be able to specify whether you want only single consonants and single vowels or whether you want to allow the possibility of combinations.

Each name type in the Simple Interface uses its own set of predefined templates, hence the diversity of names you can get from it, especially when using the "Default Names" choice. In the Advanced Interface, you can define your own by entering it into the "Name Generation Field."

In a template, a single case-sensitive character is used to represent each component of the name. 's' means syllable; 'v' means single vowel; 'V' means vowel or vowel combination. The consonants are a little more complicated, because the template distinguishes between consonant combinations that are appropriate for the beginning of names, like 'wh' and 'chr', and consonant combinations that are appropriate for the middle or ending of names, like 'nd' and 'ck'. To avoid generating names that, for example, start with 'nd' or 'ck', this distinction is necessary, although most consonant combinations are members of both groups, as they can appear anywhere. In templates, the character 'c' represents a single consonant and is appropriate anywhere in a name. The character 'B' (note that it is capital) represents a consonant or consonant combination appropriate for the beginning of a name; the character 'C' represents a consonant or consonant combination appropriate anywhere else. The apostrophe character can be used to insert an apostrophe into the name; the hyphen character can be used to insert a hyphen into the name. These symbols and other, special-purpose symbols, are all documented in the quick reference guide. Some example templates:

* `ss`
* `BVC`
* `s'vCv`
* `Bv-s`

Templates can be more complicated when you add parentheses into them. Any character in parentheses appears in the generated names literally. The template (str)VC, for example, will generate short names starting with 'str', like "Stript", "Stroagh", and "Strun". If you include an or bar, the '|' character, then either what's on the left or what's on the right will be chosen to be included in the name. The template (st|tr)VC will generate short names that start with either 'st' or 'tr'. Parentheses are useful when you already have an idea of what you want a name to look like but want to see variations on it. Here are some more sample templates:

* `ss(ien|ian)`
* `(d|t|p|w)(r)VCv`
* `ss(ly|ily|ish|ing)`

Angle brackets are another valid construct in templates. Angle brackets can be used with '|' characters in the same way as parentheses can, but what goes inside angle brackets isn't literal text but template codes for syllables, consonants, vowels, and so on. The template <s|ss>v generates names that consist either of one or two syllables, followed by a vowel. Here are some examples:

* `<c|cvc>s`
* `(ma)<VC|s|Vs>`
* `<Bv|V>(ck|kk|k)v`

If there is nothing on one side of an or bar, it allows the possibility of nothing. For example, the template (s|)(mor)ss generates names that start with 'mor' and names that start with 'smor'. The template ss<|v> generates two syllable names that may or may not have an extra vowel appended. Also, parentheses and angle brackets can be nested; the template (ghoa|<s>)VC generates names that start with either "ghoa" or a predefined syllable. Here are more examples:

* `(s||t)VC`
* `<|s|ss|sss>s`
* `<V|VCV|((a|e)i)>`

Nesting can influence the frequency certain choices are made. For example, a template starting with (r|s|t) will generate names that start with either 'r', 's', or 't', where each of the three letters appear with approximately equal frequency. But if you use a template starting with (r|(s|t)), roughly half the generated names will start with 'r', and only a quarter will start with 's', and a quarter will start with 't'.

Although the vast majority of the time, you won't need to use very complex constructions with nested parentheses and angle brackets, once in a while this functionality can be very helpful.
