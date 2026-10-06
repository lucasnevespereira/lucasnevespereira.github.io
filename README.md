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

## Files

- `static/` is the place for images and anything else that needs a short URL. inkssg copies it to the root of `public/` as is, so `static/picture.jpg` is served at `lucasnp.dev/picture.jpg`.
- `assets/` holds the favicons the theme expects under `/assets/icons/`.

The snowz page lives in `static/` as `snowz.html`, served at `lucasnp.dev/snowz`. It is not rendered by inkssg, it only reuses the theme's stylesheet.

## Deploy

Push to `main`. The `Deploy` workflow builds the site and publishes `public/` to GitHub Pages.
