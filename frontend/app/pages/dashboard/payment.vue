<template>
  <div>
    <div class="w-min text-primary-800 screen-title">
      <h1 class="text-2xl font-bold text-primary-700">Payment</h1>
      <hr></hr>
    </div>

    <ClientOnly>
      <div v-if="fetchStatus === 'pending'" class="flex justify-center items-center text-yellow-600">
        <UIcon name="i-material-symbols-hourglass-empty" class="size-5 inline-block align-middle mr-2 animate-spin" />
        <span class="align-middle">Loading...</span>
      </div>
      <div v-else class="controls my-2 p-3 border-lg border border-neutral-100 rounded-xl">
        <div class="flex gap-4 flex-wrap">
          <div class="flex gap-2 items-center">
            <UFormField label="Status">
              <USelect v-model="status" :items="filterStatus" class="w-48" />
            </UFormField>
          </div>
          <div class="flex gap-2 items-center">
            <UFormField label="Sort by">
              <USelect v-model="sortBy" :items="filterSortBy" class="w-48" />
            </UFormField>
          </div>

          <div class="flex gap-2 items-center">
            <UFormField label="Order" class="m-0 p-0">
              <USelect v-model="orderBy" :items="filterOrderBy" class="w-48" />
            </UFormField>
          </div>
        </div>
      </div>
    </ClientOnly>
    <div class="screen-content mt-4">
      <UTable
        ref="table"
        v-model:pagination="pagination"
        :data="paginationData"
        :columns="columns"

        class="flex-1"
      />

      <div class="flex justify-center border-t border-default pt-4">
        <UPagination
          :default-page="1"
          :items-per-page="pagination?.pageSize || 10"
          :total="paginationMeta?.total"
          @update:page="(p) => updatePagination(p)"
        />
      </div>
    </div>

    <!-- Modal section -->
    <UModal :title="`Reviewing Payment ${reviewRow?.getValue('payment_id')}`" v-model:open="showReviewModal">
      <template #body>
        <p class="mb-1">You are reviewing payment <strong>#{{ reviewRow?.getValue('payment_id') }}</strong>.</p>

        <p><strong>Merchant:</strong> {{ reviewRow?.getValue('merchant_name') }}</p>
        <p><strong>Amount:</strong> {{ reviewRow?.getValue('amount') }}</p>


      </template>

      <template #footer>
        <p>Please select action:</p>
        <UButton color="success" @click="() => confirmReview('completed')">Approve</UButton>
        <UButton color="error" @click="() => confirmReview('failed')">Reject</UButton>
      </template>
    </UModal>
  </div>
</template>

<script setup lang="ts">
import type { PaginationState, Row } from '@tanstack/vue-table'
import type { TableColumn } from '@nuxt/ui'
import type { BaseResponse } from '~~/types/http'
import { h, resolveComponent } from 'vue'
const UBadge = resolveComponent('UBadge')

// const table = useTemplateRef('table')

definePageMeta({
  middleware: ['auth-logged-in']
})

type Payment = {
  payment_id: string
  merchant_name: string
  date: string
  amount: string
  status: 'completed' | 'processing' | 'failed'
}

const orderBy = ref<'asc' | 'desc'>('desc')
const sortBy = ref<'date' | 'amount' | 'status'>('date')
const status = ref<'all' | 'completed' | 'processing' | 'failed'>('all')

const filterStatus = computed(() => ([
  { label: `All (${paginationExtra.value.all})`, value: 'all' },
  { label: `Completed (${paginationExtra.value.completed})`, value: 'completed' },
  { label: `Processing (${paginationExtra.value.processing})`, value: 'processing' },
  { label: `Failed (${paginationExtra.value.failed})`, value: 'failed' },
]))

const filterSortBy = [
  { label: 'Creation Date', value: 'date' },
  { label: 'Amount', value: 'amount' },
  { label: 'Status', value: 'status' },
]

const filterOrderBy = [
  { label: 'Descending', value: 'desc' },
  { label: 'Ascending', value: 'asc' },
]

watch([status, sortBy, orderBy], () => {
  // Reset to first page
  pagination.value!.pageIndex = 0
  execute()
})

const columns: TableColumn<Payment>[] = [
  {
    header: 'Payment ID',
    accessorKey: 'payment_id'
  },
  {
    header: 'Merchant',
    accessorKey: 'merchant_name'
  },
  {
    header: 'Creation Date',
    accessorKey: 'date',
    cell: ({ row }) => {
      return new Date(row.getValue('date')).toLocaleString('en-US', {
        year: 'numeric',
        day: 'numeric',
        month: 'short',
        hour: '2-digit',
        minute: '2-digit',
        hour12: false
      })
    }
  },
  {
    header: 'Amount',
    accessorKey: 'amount',
    cell: ({ row }) => {
      const amount = Number.parseFloat(row.getValue('amount'))
      const formatted = new Intl.NumberFormat('id-ID', {
        style: 'currency',
        currency: 'IDR'
      }).format(amount)
      return h('div', { class: 'text-right font-semibold' }, formatted)
    }
  },
  {
    header: 'Status',
    accessorKey: 'status',
    cell: ({ row }) => {
      const color = {
        completed: 'success' as const,
        failed: 'error' as const,
        processing: 'warning' as const
      }[row.getValue('status') as string]

      return h(UBadge, {
        class: `capitalize ${row.getValue('status') === 'processing' ? 'cursor-pointer' : ''}`,
        variant: 'subtle',
        color,
        'trailing-icon': row.getValue('status') === 'processing' ? 'i-material-symbols-arrow-right-alt': undefined,
        onClick: () => reviewEntry(row),
      }, () =>
        row.getValue('status')
      )
    },

  }
]

const pagination = ref<PaginationState>({
  pageIndex: 0,
  pageSize: 10
})

const paginationMeta = ref<BaseResponse['meta']>({
  offset: 0,
  limit: 10,
  total: 0,
})
const paginationExtra = ref({
  all: 0,
  completed: 0,
  failed: 0,
  processing: 0,
})

const toast = useToast()
const api = useApi()

function updatePagination(page: number) {
  pagination.value!.pageIndex = page
  execute()
}

// Review section
const showReviewModal = ref(false)
const reviewRow = ref<Row<Payment>>()
function reviewEntry(row: Row<Payment>) {
  if (row.getValue('status') === 'processing') {
    reviewRow.value = row
    // Open modal to review payment
    showReviewModal.value = true
    // For now, just log to console
    console.log('Reviewing payment:', row.getValue('payment_id'))
  }
}
async function confirmReview(action: 'completed' | 'failed') {
  if (!reviewRow.value) return
  const paymentId = reviewRow.value.getValue('payment_id')
  console.log(`Payment ${paymentId} marked as ${action}`)
  try {
    const response = await api<BaseResponse<{}>>(`/dashboard/v1/payment/${paymentId}/review`, {
      method: 'PUT',
      body: {
        action
      }
    })
    toast.add({
      title: 'Success',
      icon: 'i-material-symbols-check-circle-outline',
      color: 'success',
      description: (response.data as any).message || `Payment reviewed successfully to ${action}.`
    })
    // Refresh data
    execute()
  } catch (err) {
    console.error('Error reviewing payment:', err)
    toast.add({
      title: 'Error',
      icon: 'i-material-symbols-error-outline',
      color: 'error',
      description: (err as any).data?.message || 'An error occurred while reviewing the payment.'
    })
  } finally {
    // Close modal
    showReviewModal.value = false
    reviewRow.value = undefined
  }
}

const { data: paginationData, execute, status: fetchStatus } = useLazyAsyncData('payments', async () => {
  // Simulate fetching data from an API
  const response = await api<BaseResponse<Payment[]>>('/dashboard/v1/payments',{
    method: 'GET',
    params: {
      page: (pagination.value!.pageIndex || 0),
      limit: pagination.value!.pageSize || 10,
      order: orderBy.value,
      sort: sortBy.value,
      ...(status.value !== 'all' ? { status: status.value } : {}),
    }
  })
  paginationMeta.value = {
    total: response.meta?.total,
    limit: response.meta?.limit,
    offset: response.meta?.offset,
  }

  paginationExtra.value = {
    all: response.meta?.extra?.all || 0,
    completed: response.meta?.extra?.completed || 0,
    failed: response.meta?.extra?.failed || 0,
    processing: response.meta?.extra?.processing || 0,
  }


  return response.data
}, { server: false })

onMounted(() => {
  execute()
})

</script>

<style scoped>

</style>
