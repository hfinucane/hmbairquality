<script>
	import { getContext } from 'svelte';
	import ModalContent from './ModalContent.svelte'
	
	const { open } = getContext('simple-modal');

        var current;

        window.backend.NewAirQuality().then((result) => {
            current = result;
            console.log(result);
            result.results.map((x) => { console.log(x.DEVICE_LOCATIONTYPE) });
            let outside = result.results.filter((x) => { x.DEVICE_LOCATIONTYPE === "outside" });
            console.log(outside);
        });
</script>

<main>
	<p><button on:click={handleOpenModal}>Hello</button></p>
</main>

<style></style>
