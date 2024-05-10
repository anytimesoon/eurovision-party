<script lang="ts">

    import IntersectionObserver from "$lib/components/images/IntersectionObserver.svelte";
    import Image from "$lib/components/images/Image.svelte";
    import Video from "$lib/components/images/Video.svelte";

    export let src:string
    export let alt:string
    export let customClasses:string
    let intersecting = false

</script>

<IntersectionObserver once={true} bind:intersecting={intersecting}>
    {#if intersecting}
        <Video {alt} {src} {customClasses}/>
    {/if}
</IntersectionObserver>