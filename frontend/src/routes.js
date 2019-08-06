import Photos from "pages/photos.vue";
import Home from "pages/home.vue";
import Albums from "pages/albums.vue";
import Places from "pages/places.vue";
import Labels from "pages/labels.vue";
import Events from "pages/events.vue";
import People from "pages/people.vue";
import Library from "pages/library.vue";
import Share from "pages/share.vue";
import Settings from "pages/settings.vue";
import Todo from "pages/todo.vue";

export default [
    {
        name: "Home",
        path: "/",
    },
    {
        name: "Home",
        path: "/home",
        component: Home,
    },
    {
        name: "Photos",
        path: "/photos",
        component: Photos,
        meta: {area: "Photos"},
    },
    {
        name: "Albums",
        path: "/albums",
        component: Albums,
        meta: {area: "Albums"},
    },
    {
        name: "Favorites",
        path: "/favorites",
        component: Photos,
        meta: {area: "Favorites"},
        props: {staticFilter: {favorites: true}},
    },
    {
        name: "Places",
        path: "/places",
        component: Places,
        meta: {area: "Places"},
    },
    {
        name: "Labels",
        path: "/labels",
        component: Labels,
        meta: {area: "Labels"},
    },
    {
        name: "Events",
        path: "/events",
        component: Events,
        meta: {area: "Events"},
    },
    {
        name: "People",
        path: "/people",
        component: People,
        meta: {area: "People"},
    },
    {
        name: "Filters",
        path: "/filters",
        component: Todo,
        meta: {area: "Filters"},
    },
    {
        name: "Library",
        path: "/library",
        component: Library,
        meta: {area: "Library"},
    },
    {
        name: "Share",
        path: "/share",
        component: Share,
        meta: {area: "Share"},
    },
    {
        name: "Settings",
        path: "/settings",
        component: Settings,
        meta: {area: "Settings"},
    },
    {
        path: "*", redirect: "/login",
    },
];
