<template>
	<div class="flex flex-col items-center">
		<h1 class="text-gray-800 text-5xl font-bold">URL TINY</h1>
		<button
			@click="
				this.$toast.success(`Hey! I'm here`, {
					position: 'top',
				})
			"
		>
			Click for Notification
		</button>
		<transition
			enter-active-class="animate__animated animate__fadeIn"
			leave-active-class="animate__animated animate__fadeOut"
			v-on:after-leave="done = true"
			appear
		>
			<div
				class="container mx-auto flex flex-col items-center max-w-4xl"
				v-if="show"
			>
				<form
					class="shadow-xl mt-5 w-80 md:w-8/12 p-4 flex flex-col bg-white rounded-lg space-y-5 md:space-y-8"
					@submit.prevent
				>
					<label
						for="url"
						class="font-semibold text-xl"
						>Enter a URL to make TINY</label
					>
					<input
						ref="firstinput"
						v-model="url"
						name="url"
						type="url"
						autocomplete="off"
						placeholder="https://example.com/my-long-url"
						class="mb-3 py-3 px-4 border border-gray-400 focus:outline-none rounded-md focus:ring-1 ring-cyan-500"
					/>
					<div
						class="flex items-center space-x-2 w-full"
					>
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
		</transition>
		<transition
			enter-active-class="animate__animated animate__fadeIn"
			leave-active-class="animate__animated animate__fadeOut"
			appear
		>
			<div
				class="container mx-auto flex flex-col items-center max-w-4xl"
				v-if="done"
			>
				<div
					class="shadow-xl mt-5 w-80 md:w-8/12 p-4 flex flex-col bg-white rounded-lg space-y-5 md:space-y-8"
				>
					<p class="font-semibold text-xl">
						Your generated URL is
					</p>
					<div
						class="flex items-center space-x-2 w-full"
					>
						<input
							type="text"
							:disabled="true"
							ref="text"
							v-model="genurl"
							class="py-3 px-4 border border-gray-400 rounded-md w-full font-semibold text-md"
						/>
						<button
							class="py-3 px-3 border border-gray-400 rounded-md"
							@click="copy"
						>
							<svg
								class="h-6 w-6"
								xmlns="http://www.w3.org/2000/svg"
								fill="none"
								viewBox="0 0 24 24"
								stroke="currentColor"
							>
								<path
									strokeLinecap="round"
									strokeLinejoin="round"
									strokeWidth="{2}"
									d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
								/>
							</svg>
						</button>
					</div>
				</div>
			</div>
		</transition>
	</div>
</template>

<script>
export default {
	name: "UrlForm",
	data() {
		return {
			genurl: "",
			url: "",
			done: false,
			short: "",
			show: true,
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
					if (res.status == 200) {
						this.genurl = `urltiny.in/${this.short}`;
						this.show = false;
					} else {
						this.url = "";
						this.short = "";
						this.$refs.firstinput.focus();
						this.$toast.error(
							`SHORT URL already in use`,
							{ position: "top" }
						);
					}
					return res.json();
				})
				.then((res) => {
					/* console.log(res); */
				})
				.catch((err) => {
					console.error(err.messsage);
				});
		},
		copy: function () {
			console.log("test");
			this.selectText(this.$refs.text);
			document.execCommand("copy");
		},
		selectText(element) {
			var range;
			if (document.selection) {
				// IE
				range = document.body.createTextRange();
				range.moveToElementText(element);
				range.select();
			} else if (window.getSelection) {
				range = document.createRange();
				range.selectNode(element);
				window.getSelection().removeAllRanges();
				window.getSelection().addRange(range);
			}
		},
	},
};
</script>
