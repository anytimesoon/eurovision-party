<script lang="ts">
    import { onMount } from 'svelte'

    interface Props {
        src: string;
        customClasses: string;
    }

    let { src, customClasses }: Props = $props();
    let loaded = $state(false)
    let thisImage:HTMLVideoElement = $state()

    onMount(() => {
        thisImage.onloadeddata = () => {
            console.log('loaded')
            loaded = true
        }
    })

    let imageOpacity = $derived(loaded ? "opacity-100" : "opacity-0")
    let pulse = $derived(loaded ? "" : "bg-canvas-secondary animate-pulse")
</script>

<div class="{pulse}">
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
    {:else}
        <div class="bg-canvas-secondary">
            <video autoplay loop muted playsinline bind:this={thisImage} class="{imageOpacity} {customClasses} transition-opacity ease-out duration-1000" >
                <source src={src} type="video/*" />
            </video>
        </div>
    {/if}
</div>