<script lang="ts">

    import IntersectionObserver from "$lib/components/images/IntersectionObserver.svelte";
    import Video from "$lib/components/images/Video.svelte";

    export let src:string
    export let customClasses:string
    let intersecting = false

</script>

<IntersectionObserver once={true} bind:intersecting={intersecting}>
    {#if intersecting}
        <Video {src} {customClasses}/>
    {/if}
</IntersectionObserver>