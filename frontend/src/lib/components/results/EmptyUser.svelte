<script lang="ts">
    import {userStore} from "$lib/stores/user.store";
    import {resultPageState} from "$lib/stores/resultPageState.store";
    import SelectUserResultForm from "$lib/components/forms/SelectUserResultForm.svelte";
    import ResetUserResultForm from "$lib/components/forms/ResetUserResultForm.svelte";
</script>

<div class="text-center">
    <h1 class="font-sans">🤷</h1>
    <p>
        Looks like {$userStore.get($resultPageState.userId).name} hasn't voted yet.
    </p>
    <SelectUserResultForm />
    <ResetUserResultForm />
</div>