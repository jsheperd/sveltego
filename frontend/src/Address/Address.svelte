<script>

    let count = 1
    let src = "/hello"

    function PushMe() {
        console.log("I was pushed !")
        count += 1
    }

    async function FecthHello() {
		console.log(src)
        const res = await fetch(src);
        const json = await res.json();
        console.log(json);
    }

</script>

<main>
	<h1>{src}</h1>
	<h2>Hall</h2>
    <button on:click="{PushMe}"> Push me {count}</button>
    <input type="text" name="src" bind:value={src} />
    <button on:click="{FecthHello}"> Fetch </button>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #00ff22;
		text-transform: uppercase;
		font-size: 8em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>