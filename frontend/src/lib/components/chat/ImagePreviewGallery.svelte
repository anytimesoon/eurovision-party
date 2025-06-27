<script lang="ts">
    import {quintInOut} from "svelte/easing";
    import { scale } from 'svelte/transition';
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";

    interface Props {
        previewImage: string;
        fileName?: string;
        cancelUpload: VoidFunction;
    }

    let { previewImage, fileName = "", cancelUpload }: Props = $props();
</script>

{#if fileName !== ``}
    <div transition:scale|global={{ duration: 500, opacity: 0.5, easing: quintInOut }} class="flex">
        <div class="h-[60px] overflow-hidden rounded mb-2">
            <img src={previewImage} alt="preview" class="h-[60px]"/>
        </div>
        <div class="pl-3">
            <button onclick={cancelUpload}>
                <CloseCircleOutline />
            </button>
        </div>
    </div>
{/if}