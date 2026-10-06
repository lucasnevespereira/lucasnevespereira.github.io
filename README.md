# lucasnevespereira.github.io

Personal site at [lucasnp.dev](https://lucasnp.dev), built with [inkssg](https://github.com/snowztech/inkssg).

## Build

```
make site
```

Renders `pages/index/content.md` with the `minimal` theme into `public/`. Site data lives in `ink.yaml`.

## Preview

```
python3 -m http.server 8000 --directory public
```

- Main site: http://localhost:8000
- snowz page: http://localhost:8000/snowz.html

The snowz page is a standalone HTML file at `snowz.html`. It is not rendered by inkssg, `make site` copies it and `snowz/` into `public/`. It reuses the theme's stylesheet.

## Deploy

Push to `main`. The `Deploy` workflow builds the site and publishes `public/` to GitHub Pages.
