<script lang="ts">

    import IntersectionObserver from "$lib/components/images/IntersectionObserver.svelte";
    import Image from "$lib/components/images/Image.svelte";

    interface Props {
        src: string;
        alt: string;
        customClasses: string;
    }

    let { src, alt, customClasses }: Props = $props();
    let intersecting = $state(false)

</script>

<IntersectionObserver once={true} bind:intersecting={intersecting}>
    {#if intersecting}
        <Image {alt} {src} {customClasses}/>
    {/if}
</IntersectionObserver>