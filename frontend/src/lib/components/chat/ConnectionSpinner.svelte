<script lang="ts">
    import Spinner from "$lib/components/Spinner.svelte";
    import {socketStateStore} from "$lib/stores/socketState.store";
    import {socketRetryCount} from "$lib/stores/socketRetryCount.store";
</script>

<div class="relative">
    {#if $socketStateStore === false && $socketRetryCount > 3}
        <div class="mx-auto w-[50%] z-40 absolute top-5 left-0 right-0 bottom-0">
            <div class="bg-gray-600 rounded py-3 shadow-lg shadow-gray-800">
                <p class="text-center">
                    Connecting to the chat
                    <span class="pl-1">
                    <Spinner size={'s'} thickness={'s'} color={'grey'} isBlock={false}/>
                </span>
                </p>
            </div>
        </div>
    {/if}
</div>
