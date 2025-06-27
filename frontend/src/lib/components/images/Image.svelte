<script lang="ts">
    import { onMount } from 'svelte'
    import {sessionStore} from "$lib/stores/session.store";

    interface Props {
        src: string;
        alt: string;
        customClasses: string;
    }

    let { src, alt, customClasses }: Props = $props();
    let loaded = $state(false)
    let thisImage:HTMLImageElement = $state()
    let imageBlobUrl:string = $state("")

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

    let imageOpacity = $derived(loaded ? "opacity-100" : "opacity-0")
    let pulse = $derived(loaded ? "" : "bg-canvas-secondary animate-pulse")
    $effect(() => {
        if (src) {
            refresh()
        }
    });
</script>

<div class="{pulse} h-full w-full">
    <img src={imageBlobUrl} {alt} class="{imageOpacity} {customClasses} transition-opacity ease-out duration-1000" bind:this={thisImage} />
</div>
