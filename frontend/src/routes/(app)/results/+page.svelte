<script lang="ts">
    import {onMount} from "svelte";
    import {results} from "$lib/stores/results.store";
    import ResultsTable from "$lib/components/results/ResultsTable.svelte";

    onMount(async () => {
        await results.refresh()
    })
</script>

<div class="h-full flex flex-col">

    <h1 class="text-center">Rankings</h1>

    <div class="flex-1 overflow-auto">
        <div class="rounded max-h-1">
            <div class="py-3">

                <ResultsTable />

            </div>
        </div>
    </div>
</div>
