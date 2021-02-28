<template>
	<div class="flex flex-col items-center">
		<h1 class="text-gray-800 text-5xl font-bold">URL TINY</h1>
		<div
			class="container mx-auto flex flex-col items-center max-w-4xl"
		>
			<form
				class="shadow-xl mt-5 w-80 md:w-8/12 p-4 flex flex-col bg-white rounded-lg space-y-5 md:space-y-8"
			>
				<label for="url" class="font-semibold text-xl"
					>Enter a URL to make TINY</label
				>
				<input
					v-model="url"
					name="url"
					type="url"
					autocomplete="off"
					placeholder="https://example.com/my-long-url"
					class="mb-3 py-3 px-4 border border-gray-400 focus:outline-none rounded-md focus:ring-1 ring-cyan-500"
				/>
				<div class="flex items-center space-x-2 w-full">
					<label
						for="short"
						class="py-3 px-4 font-semibold text-md border border-gray-400 rounded-md"
						>urltiny.in/</label
					>
					<input
						v-model="short"
						name="short"
						autocomplete="off"
						type="text"
						placeholder="shorturl"
						class="py-3 px-4 w-full border border-gray-400 focus:outline-none rounded-md focus:ring-1 ring-cyan-500"
					/>
				</div>
				<button
					@click="newUrl()"
					class="w-full bg-blue-500 text-white p-3 rounded-lg font-semibold text-lg"
				>
					Generate
				</button>
			</form>
		</div>
	</div>
</template>

<script>
export default {
	name: "UrlForm",

	data() {
		return {
			url: "",
			short: "",
		};
	},
	methods: {
		newUrl: function () {
			fetch("https://api.urltiny.in/newurl", {
				method: "POST",
				headers: {
					Accept:
						"application/json, text,plain, */*",
					"Content-Type": "application/json",
				},
				body: JSON.stringify({
					url: this.url,
					short: this.short,
				}),
			})
				.then((res) => {
					return res.json();
				})
				.then((res) => {
					console.log(res);
				})
				.catch((err) => {
					console.error(err.messsage);
				});
		},
	},
};
</script>
