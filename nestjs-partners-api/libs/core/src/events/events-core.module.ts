import { Module } from '@nestjs/common';
import { EventsService } from './events.service';
import { PrismaService } from '../prisma/prisma.service';
import { PrismaModule } from '../prisma/prisma.module';

@Module({
  imports:[PrismaModule],
  providers: [EventsService],
  exports: [EventsService],
})
export class EventsCoreModule {}
