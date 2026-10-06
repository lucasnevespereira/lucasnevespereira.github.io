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

- `assets/` holds icons and images. inkssg copies it to `public/assets/`, so `assets/img/picture.jpg` is served at `lucasnp.dev/assets/img/picture.jpg`.
- `static/` holds files that need a root URL. inkssg copies it to the root of `public/` as is.

The snowz page lives in `static/` as `snowz.html`, served at `lucasnp.dev/snowz`. It is not rendered by inkssg, it only reuses the theme's stylesheet.

## Deploy

Push to `main`. The `Deploy` workflow builds the site and publishes `public/` to GitHub Pages.
