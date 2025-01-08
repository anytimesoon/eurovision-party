<script lang="ts">
    import Refresh from "svelte-material-icons/Refresh.svelte";
    import {resultPageState} from "$lib/stores/resultPageState.store";
    import {results} from "$lib/stores/results.store";

    const resetFilters = async () => {
        resultPageState.reset()
        results.refresh()
    }
</script>



<button class="bg-canvas-primary border-2 border-primary pl-2" on:click={resetFilters}>
    <span class="flex text-sm"><Refresh size="1.3em" class="pt-0.5"/> Reset</span>
</button>


