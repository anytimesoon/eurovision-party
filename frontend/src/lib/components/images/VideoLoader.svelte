<script lang="ts">

    import IntersectionObserver from "$lib/components/images/IntersectionObserver.svelte";
    import Video from "$lib/components/images/Video.svelte";

    interface Props {
        src: string;
        customClasses: string;
    }

    let { src, customClasses }: Props = $props();
    let intersecting = $state(false)

</script>

<IntersectionObserver once={true} bind:intersecting={intersecting}>
    {#if intersecting}
        <Video {src} {customClasses}/>
    {/if}
</IntersectionObserver>