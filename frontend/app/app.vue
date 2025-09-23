<template>
  <UApp :toaster="{ position: 'top-right', duration: 3000 }">
    <div class="grid grid-cols-[200px_1fr] grid-rows-[64px_1fr] gap-8">
      <div class="col-span-2 row-span-1 p-4 flex gap-2 items-center bg-gray-50 shadow-md rounded-b-xl relative">
        <UIcon name="i-material-symbols-hexagon" class="size-12 text-red-500" />
        <p class="text-lg font-semibold">Durianpay CS Dashboard</p>
        <div class="logout-section right-0 absolute m-2 flex gap-4 items-center" v-if="store.isLoggedIn">
          <p class="bg-white text-primary-500 px-3 py-1 rounded-xl">Welcome, <span class="font-semibold">{{ store.user?.username }}!</span></p>
          <UButton @click="store.logout()" :disabled="store.loading" class="mr-2">
            <UIcon name="i-material-symbols-logout" class="size-5 inline-block align-middle" />
            <span class="align-middle">Logout</span>
          </UButton>
        </div>
      </div>
      <div class="col-span-1 row-span-1">
        <nav class="py-4 rounded-r-2xl bg-green-50">
          <ul>
            <NuxtLink to="/dashboard/payment"><DashboardMenu>Payment</DashboardMenu></NuxtLink>
            <DashboardMenu disabled>Customers</DashboardMenu>
            <DashboardMenu disabled>Subscriptions</DashboardMenu>
            <DashboardMenu disabled>Products</DashboardMenu>
            <hr class="m-4 text-green-300"></hr>
            <DashboardMenu disabled>Users</DashboardMenu>
          </ul>
        </nav>
      </div>
      <div class="col-span-1 row-span-1 m-4 min-h-96 relative">
        <NuxtPage />
        <div class="footer w-full max-w-sm mt-8">
          <hr class="border-t border-green-300"></hr>
          <p class="text-sm text-green-500 mt-2">© 2025 Durianpay</p>
        </div>
      </div>
    </div>
  </UApp>
</template>

<script setup lang="ts">
import { useAuthStore } from './stores/auth';

const store = useAuthStore();
</script>
