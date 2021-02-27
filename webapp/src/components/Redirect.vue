<template>
	{{ $route.params.short }}
</template>

<script>
export default {
	created() {
		/* window.location.href = "https://duckduckgo.com" */
		fetch(
			`http://localhost:3000/urls?short=${this.$route.params.short}`
		)
			.then((res) => {
				return res.json();
			})
			.then((res) => {
				/* console.log(res); */

				window.location.href = res.url;
			});
	},
};
</script>

<style></style>
