import { test, expect } from '@playwright/test'

test.describe('Test navbar', () => {
  test('navigation works', async ({ page }) => {
    await page.goto('/')

    // Click navbar link "Test"
    await page.locator('a >> text=Journeys').click()
    await expect(page).toHaveURL('/journeys')

    // Click navbar link "Home"
    await page.locator('a >> text=Home').click()
    await expect(page).toHaveURL('/')
  })
})
