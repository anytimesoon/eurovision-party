<script lang="ts">
    import { onMount } from 'svelte'

    export let src:string
    export let alt:string
    export let customClasses:string
    let loaded = false
    let thisImage:HTMLImageElement

    onMount(() => {
        thisImage.onload = () => {
            loaded = true
        }
    })

    $: imageOpacity = loaded ? "opacity-100" : "opacity-0"
</script>

<img {src} {alt} class="{imageOpacity} {customClasses} transition-opacity ease-out duration-1000" bind:this={thisImage} />