# GoSSaurus - Yet another Kindle dictionary Generator

### generate a rich dictionary for your kindle

### dependencies: kindlegen_linux

## Usage
```bash
 gossaurus build 
    --from=hr \
    --to=en \
    --title="Serbian to English Dictionary" \
    --author="Milos" \
    --cover="cover.png" \
    --entries=definitions.tsv
```
### --from and --to
* Language codes are described as [ISO 639-1 Codes](https://en.wikipedia.org/wiki/List_of_ISO_639_language_codes)
* Some languages might use alternative codes to default to similar, more popular languages, and they are listed below
<details>
  <summary>Default Language Codes</summary>
| Language | Code |
|:-------- | :---: |
| Serbo-Croation | hr |
</details>

### --entries=definitions.tsv
* see [example.tsv](docs/example.tsv)

---

## Install and Config (TODO)
```aiignore
make install
```

## Building from source (TODO)
```aiignore
make build
```

## FAQ - Troubleshooting
1. My dictionary is not showing definitions in my book
   Check that the language codes in your book and in your dictionary (from) are the same.
   Some language codes are not correctly defined, and you might have to adjust to the correct code in your book.
   You can use calibre for that, for example.

Command below will fix the language metadata of your kindle book (if it's not locked)
```
ebook-convert neuromancer.mobi neuromancer-serbian.mobi --language sr
```

Even if you set "sr" for serbian, it will set the language to "hr" (croatian).
Some languages are tricky (why?)
What about serbian cirilics, is there a distiction for that? (#todo)


