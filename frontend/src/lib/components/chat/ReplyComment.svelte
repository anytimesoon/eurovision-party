<script lang="ts">
    import {quintInOut} from "svelte/easing";
    import {replyComment} from "$lib/stores/replyComment.store";
    import {userStore} from "$lib/stores/user.store";
    import CloseCircleOutline from "svelte-material-icons/CloseCircleOutline.svelte";
    import { scale } from 'svelte/transition';
    import {staticEP} from "$lib/models/enums/endpoints.enum";
    import ImageLoader from "$lib/components/images/ImageLoader.svelte";

    let shouldDisplay = $state(false)
    let fileName = $state("")

    function close() {
        shouldDisplay = false
        replyComment.close()
    }

    $effect(() => {
        if($replyComment) {
            shouldDisplay = $replyComment.text !== "" || $replyComment.fileName !== ""
            fileName = $replyComment.fileName
        }
    });

</script>

{#if shouldDisplay}
    <div transition:scale|global={{ duration: 500, opacity: 0.5, easing: quintInOut }} class="bg-canvas-primary p-2 mb-2 rounded text-typography-main text-xs relative">
        <button class="bg-transparent absolute top-1 right-1"  onclick={close}>
            <CloseCircleOutline />
        </button>

        {#if $replyComment.userId !== undefined && $userStore.get($replyComment.userId) !== undefined}
            <div class="pb-2">
                {$userStore.get($replyComment.userId).name}
            </div>
        {/if}
        <div class="flex">
            {#if $replyComment.fileName !== ""}
                <div class="mr-3 max-h-[20px] rounded overflow-hidden">
                    <ImageLoader src={staticEP.CHAT_IMG + fileName} alt="comment image" customClasses="h-[20px]"/>
                </div>
            {/if}
            <div class="flex-1">
                {$replyComment.text}
            </div>
        </div>
    </div>
{/if}