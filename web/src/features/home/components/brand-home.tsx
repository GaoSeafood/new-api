/*
Copyright (C) 2023-2026 QuantumNous

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <https://www.gnu.org/licenses/>.

For commercial licensing, please contact support@quantumnous.com
*/
import { Link } from '@tanstack/react-router'
import { ArrowRight, Cable, ChartNoAxesCombined, KeyRound } from 'lucide-react'
import { useTranslation } from 'react-i18next'

import brandMascot from '@/assets/brand-mascot.png'
import { Button } from '@/components/ui/button'
import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card'
import { useStatus } from '@/hooks/use-status'

interface BrandHomeProps {
  isAuthenticated: boolean
}

export function BrandHome(props: BrandHomeProps) {
  const { t } = useTranslation()
  const { status } = useStatus()
  const docsUrl = (status?.docs_link as string | undefined) || '/docs'
  const docsIsExternal = docsUrl.startsWith('http')

  const workflowSteps = [
    {
      number: '01',
      title: t('Connect models'),
      description: t(
        'Route requests through the models configured for your service.'
      ),
      icon: Cable,
    },
    {
      number: '02',
      title: t('Manage access'),
      description: t(
        'Create API keys and manage their access from the console.'
      ),
      icon: KeyRound,
    },
    {
      number: '03',
      title: t('Follow usage'),
      description: t('Review request logs and usage details as you work.'),
      icon: ChartNoAxesCombined,
    },
  ]

  return (
    <main className='overflow-hidden'>
      <section className='relative isolate px-5 pt-28 pb-16 sm:px-8 md:pt-36 md:pb-24'>
        <div
          aria-hidden='true'
          className='pointer-events-none absolute inset-0 -z-10 bg-[radial-gradient(ellipse_at_18%_12%,oklch(0.96_0.035_78_/_85%),transparent_48%),radial-gradient(ellipse_at_82%_25%,oklch(0.82_0.12_42_/_28%),transparent_42%)] dark:opacity-30'
        />
        <div className='mx-auto grid max-w-7xl items-center gap-12 lg:grid-cols-[0.9fr_1.1fr] lg:gap-16'>
          <div className='max-w-2xl'>
            <div className='border-primary/20 bg-primary/5 text-primary mb-6 inline-flex items-center gap-2 rounded-full border px-3.5 py-2 text-xs font-semibold tracking-wide'>
              <span className='bg-primary size-1.5 rounded-full' />
              {t('One gateway. Your way.')}
            </div>
            <h1 className='text-4xl leading-[1.08] font-semibold tracking-[-0.045em] text-balance sm:text-5xl md:text-6xl'>
              {t('A clearer way to build with AI.')}
            </h1>
            <p className='text-muted-foreground mt-6 max-w-xl text-base leading-7 sm:text-lg sm:leading-8'>
              {t(
                'Connect the models you use through one API, then manage keys, requests and usage in one place.'
              )}
            </p>
            <div className='mt-8 flex flex-wrap items-center gap-3'>
              <Button
                className='h-11 rounded-xl px-5 text-sm font-semibold'
                render={
                  <Link
                    to={props.isAuthenticated ? '/dashboard' : '/sign-up'}
                  />
                }
              >
                {t(props.isAuthenticated ? 'Go to Dashboard' : 'Get Started')}
                <ArrowRight className='ml-1 size-4' aria-hidden='true' />
              </Button>
              <Button
                variant='outline'
                className='h-11 rounded-xl px-5 text-sm font-semibold'
                render={<Link to='/pricing' />}
              >
                {t('Explore models')}
              </Button>
              <Button
                variant='ghost'
                className='h-11 rounded-xl px-4 text-sm font-medium'
                render={
                  docsIsExternal ? (
                    <a href={docsUrl} target='_blank' rel='noreferrer' />
                  ) : (
                    <Link to={docsUrl} />
                  )
                }
              >
                {t('Docs')}
              </Button>
            </div>
          </div>

          <div className='relative mx-auto w-full max-w-[42rem]'>
            <div
              aria-hidden='true'
              className='bg-primary/15 absolute -inset-4 -z-10 rounded-[2.25rem] blur-3xl'
            />
            <div className='relative overflow-hidden rounded-[1.75rem] bg-[#120d0a] shadow-[0_30px_90px_-42px_rgba(84,39,19,0.65)] ring-1 ring-black/10'>
              <img
                src={brandMascot}
                alt=''
                aria-hidden='true'
                className='block h-auto w-full object-contain'
                width={1536}
                height={1024}
                fetchPriority='high'
              />
              <div className='absolute right-4 bottom-4 left-4 flex items-center justify-between rounded-xl border border-white/10 bg-black/45 px-4 py-3 text-white/90 backdrop-blur-md sm:right-5 sm:bottom-5 sm:left-5'>
                <span className='text-xs font-medium sm:text-sm'>
                  {t('A friendly home for your AI workflow')}
                </span>
                <span className='text-[10px] font-semibold tracking-[0.18em] text-white/55 uppercase'>
                  API · MODELS · USAGE
                </span>
              </div>
            </div>
          </div>
        </div>
      </section>

      <section className='px-5 pt-10 pb-20 sm:px-8 md:pt-14 md:pb-28'>
        <div className='mx-auto max-w-7xl'>
          <div className='mb-8 max-w-2xl md:mb-12'>
            <p className='text-primary mb-3 text-xs font-semibold tracking-[0.18em] uppercase'>
              {t('A simpler path from model to usage')}
            </p>
            <h2 className='text-2xl leading-tight font-semibold tracking-tight text-balance sm:text-3xl md:text-4xl'>
              {t('Your AI workflow, all in one clear place.')}
            </h2>
          </div>
          <div className='grid gap-4 md:grid-cols-3 md:gap-5'>
            {workflowSteps.map((step) => {
              const Icon = step.icon
              return (
                <Card
                  key={step.number}
                  className='bg-card/80 ring-border/70 hover:bg-card min-h-52 rounded-2xl border-transparent p-1 shadow-none transition-colors duration-200'
                >
                  <CardHeader className='gap-5 px-5 pt-5 sm:px-6 sm:pt-6'>
                    <div className='flex items-center justify-between'>
                      <span className='bg-primary/10 text-primary flex size-11 items-center justify-center rounded-xl'>
                        <Icon className='size-5' strokeWidth={1.8} />
                      </span>
                      <span className='text-muted-foreground/55 font-mono text-xs tabular-nums'>
                        {step.number}
                      </span>
                    </div>
                    <CardTitle className='text-lg font-semibold tracking-tight'>
                      {step.title}
                    </CardTitle>
                  </CardHeader>
                  <CardContent className='px-5 pb-5 sm:px-6 sm:pb-6'>
                    <p className='text-muted-foreground text-sm leading-6'>
                      {step.description}
                    </p>
                  </CardContent>
                </Card>
              )
            })}
          </div>
        </div>
      </section>
    </main>
  )
}
