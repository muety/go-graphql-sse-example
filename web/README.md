# go-graphql-sse-example

## Introduction
This project aims to provide a good starting point for a Vue-based SPA frontend with a clean code structure and separation of concerns. Of course, software architecture is quite subjective and a matter of taste to some extent. But the present structure has proven to be very effective in several project of ours.

### Project structure
* `|_ components`: Reusable, custom HTML elements with no dependencies on external services.
* `|_ pages`: Components that correspond to entire views within your application. Usually routable, i.e. can be addressed with a URL. Like [Activities](https://developer.android.com/reference/android/app/Activity) in Android.   
* `|_ filters`: Functional transformations to be performed on attributes and values inside HTML templates.
* `|_ mixins`: Component methods, to be commonly reused among several components or pages. Use with care and only if you really need them, as mixins are often considered an anti-pattern in Vue. Similar to [traits]([Trait](https://docs.scala-lang.org/tour/traits.html)) in Scala. 
* `|_ model`: Definitions of data classes and entities, representing classes of objects in your business domain.
* `|_ store`: Reactive Vuex stores, separated by business domain.
* `|_ styles`: Application-wide, global CSS classes.
* `|_ plugins`: Vue [plugins](https://vuejs.org/v2/guide/plugins.html), i.e. functionality to globally extend Vue with. Might also contain adapters to third-party services.
* `|_ vendor`: Any third-party code or libraries that can not be installed via NPM. Usually static functions you copy from StackOverflow, etc. 
* `App.vue`: Root component, defining your application's very basic layout.
* `main.js`: Application entry point. Usually only modified when declaring new plugins, etc.
* `router`: Your `vue-router` route definitions.

### Fundamental Guidelines
* All data is managed by the reactive Vuex [store](src/store) and data updates should only happen through mutations.
* All interactions with the backend are performed inside store actions.
* Every entity is modeled as an instance of a class, defined inside [models](src/model). Entity-specific functionality is implemented as class methods, rather than static util functions (e.g. see `get sum()` in [Cart](src/model/cart.js)). 
* Only [pages](src/pages) may ever interact with the store. [Components](src/components) receive and send all required in- and output data through props and events to their parent components.

### Limitations
* Instead of a bunch of custom-defined [CSS](src/styles/index.css) classes, you may want to use something like [Tailwind](https://tailwindcss.com) instead
* See [README](../README.md#Limitations) 

## Project setup
```
yarn install
```

### Compiles and hot-reloads for development
```
yarn serve
```

### Compiles and minifies for production
```
yarn build
```

### Lints and fixes files
```
yarn lint
```

### Customize configuration
See [Configuration Reference](https://cli.vuejs.org/config/).
