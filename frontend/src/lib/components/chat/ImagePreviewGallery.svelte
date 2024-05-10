<script lang="ts">
    import {quintInOut} from "svelte/easing";
    import { scale } from 'svelte/transition';
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";

    export let previewImage:string
    export let fileName = ""
    export let cancelUpload:VoidFunction
    let gallery:HTMLDivElement
</script>

{#if fileName !== ``}
    <div bind:this={gallery} transition:scale={{ duration: 500, opacity: 0.5, easing: quintInOut }} class="flex">
<!--        <img class="h-[60px] mb-2" src={previewImage} alt="preview"/>-->
        <div class="pl-3">
            <button on:click={cancelUpload}>
                <CloseCircleOutline />
            </button>
        </div>
    </div>
{/if}