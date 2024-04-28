<script lang="ts">
    import Refresh from "svelte-material-icons/Refresh.svelte";
    import { enhance } from '$app/forms';
    import {resultPageState} from "$lib/stores/resultPageState.store";
    import {voteCats} from "$lib/models/enums/categories.enum";

    let form:HTMLFormElement

    const resetFilters = () => {
        $resultPageState.category = voteCats.TOTAL
        form.requestSubmit()
    }
</script>


<form method="POST" action="?/getUserResults" use:enhance bind:this={form}>

    <select class="hidden" name="id">
        <option value="">Everyone</option>
    </select>

    <button class="bg-canvas-primary border-2 border-primary pl-2" on:click|preventDefault={resetFilters}>
        <span class="flex text-sm"><Refresh size="1.3em" class="pt-0.5"/> Reset</span>
    </button>
</form>


