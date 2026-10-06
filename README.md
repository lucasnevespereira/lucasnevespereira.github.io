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

## Static files

`static/` is the site root. Everything in it is copied to `public/` as is, so the path in the repo is the path on the site: `static/picture.jpg` is served at `lucasnp.dev/picture.jpg`. Icons, images and standalone HTML pages all go here.

The snowz page lives there as `static/snowz.html`. It is not rendered by inkssg, it only reuses the theme's stylesheet.

## Deploy

Push to `main`. The `Deploy` workflow builds the site and publishes `public/` to GitHub Pages.
