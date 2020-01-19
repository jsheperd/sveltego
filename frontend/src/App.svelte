<script>
	export let name;
	import Address from "./Address/Address.svelte"
</script>

<main>
	<h1>Hello very {name}!</h1>
	<p>Visit the <a href="https://svelte.dev/tutorial">Svelte  Basic tutorial</a> to learn how to build Svelte apps.</p>
	<Address />
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>