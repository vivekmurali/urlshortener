import { createWebHistory, createRouter } from "vue-router";
import UrlForm from "./components/UrlForm.vue";
import Redirect from "./components/Redirect.vue";
// import Recipes from "./components/Recipes.vue";
const history = createWebHistory();
const routes = [
	{ path: "/", component: UrlForm },
	{ path: "/:short", component: Redirect },
	// { path: "/recipes", component: Recipes },
];
const router = createRouter({ history, routes });
export default router;
