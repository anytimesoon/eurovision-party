<script>
    import { page } from '$app/stores';
</script>

<div class="h-screen w-full flex justify-center items-center">
    <div class="text-center p-3">
        <h1 class="font-sans">🫠</h1>
        <h3>The crew is partying too hard back stage.</h3>
        <p>Something went wrong with your request, please try again.</p>
        <p class="text-[0.6rem]">{$page.status}: {$page.error.message}</p>
    </div>
</div>