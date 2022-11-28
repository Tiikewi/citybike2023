import { test, expect } from '@playwright/test'

test.describe('Test navbar', () => {
  test('navigation works', async ({ page }) => {
    // 10 journeys + 1 header row
    const expectedRowAmount = 11
    await page.goto('/journeys')

    await page.locator('tr').first().waitFor()
    const count = await page.locator('tr').count()
    expect(count).toBe(expectedRowAmount)
  })
})
