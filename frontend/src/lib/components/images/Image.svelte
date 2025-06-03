<script lang="ts">
    import { onMount } from 'svelte'
    import {sessionStore} from "$lib/stores/session.store";

    export let src:string
    export let alt:string
    export let customClasses:string
    let loaded = false
    let thisImage:HTMLImageElement
    let imageBlobUrl:string = ""

    onMount(async () => {
        thisImage.onload = () => {
            loaded = true
        }

        await refresh()
    })

    async function refresh() {
        const response = await fetch(src, {
            headers: {
                Authorization: $sessionStore, // Authorization token
            },
        });
        const blob = await response.blob();
        imageBlobUrl = URL.createObjectURL(blob);
    }

    $: imageOpacity = loaded ? "opacity-100" : "opacity-0"
    $: pulse = loaded ? "" : "bg-canvas-secondary animate-pulse"
    $: if (src) {
        refresh()
    }
</script>

<div class="{pulse} h-full w-full">
    <img src={imageBlobUrl} {alt} class="{imageOpacity} {customClasses} transition-opacity ease-out duration-1000" bind:this={thisImage} />
</div>
