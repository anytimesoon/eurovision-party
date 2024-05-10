<script lang="ts">
    import { onMount } from 'svelte'

    export let src:string
    export let alt:string
    export let customClasses:string
    let loaded = false
    let thisImage:HTMLVideoElement

    onMount(() => {
        thisImage.onloadeddata = () => {
            console.log('loaded')
            loaded = true
        }
    })

    $: imageOpacity = loaded ? "opacity-100" : "opacity-0"
    $: pulse = loaded ? "" : "animate-pulse"
</script>

<div class="bg-canvas-secondary {pulse}">
    {#if src.includes("mp4")}
        <div class="bg-canvas-secondary">
            <video autoplay loop muted playsinline bind:this={thisImage} class="{imageOpacity} {customClasses} transition-opacity ease-out duration-1000" >
                <source src={src} type="video/mp4" />
            </video>
        </div>
    {:else if src.includes("webm")}
        <div class="bg-canvas-secondary">
            <video autoplay loop muted playsinline bind:this={thisImage} class="{imageOpacity} {customClasses} transition-opacity ease-out duration-1000" >
                <source src={src} type="video/webm" />
            </video>
        </div>
    {/if}
</div>