<script lang="ts">
    import {quintInOut} from "svelte/easing";
    import {replyComment} from "$lib/stores/replyComment.store";
    import {userStore} from "$lib/stores/user.store";
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";
    import { scale } from 'svelte/transition';
    import {staticSvelteEP} from "$lib/models/enums/endpoints.enum";

    let shouldDisplay = false

    function close() {
        shouldDisplay = false
        replyComment.close()
    }

    $: if($replyComment) {
        shouldDisplay = $replyComment.text !== undefined;
    }
</script>

{#if shouldDisplay}
    <div transition:scale={{ duration: 500, opacity: 0.5, easing: quintInOut }} class="bg-canvas-primary p-2 mb-2 rounded text-typography-main text-xs relative">
        <button class="bg-transparent absolute top-1 right-1"  on:click={close}>
            <CloseCircleOutline />
        </button>

        {#if $replyComment.userId !== undefined}
            <div class="pb-2">
                {$userStore[$replyComment.userId].name}
            </div>
        {/if}
        <div class="flex">
            {#if $replyComment.fileName !== ""}
                <div class="mr-3 max-h-[20px] rounded overflow-hidden">
                    <img src={staticSvelteEP.CHAT_IMG + $replyComment.fileName} alt="" class="h-[20px]"/>
                </div>
            {/if}
            <div class="flex-1">
                {$replyComment.text}
            </div>
        </div>
    </div>
{/if}