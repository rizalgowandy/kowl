/**
 * Copyright 2022 Redpanda Data, Inc.
 *
 * Use of this software is governed by the Business Source License
 * included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
 *
 * As of the Change Date specified in that file, in accordance with
 * the Business Source License, use of this software will be governed
 * by the Apache License, Version 2.0
 */

import { observer } from 'mobx-react';
import React, { Component } from 'react';

import { PencilIcon, TrashIcon } from '@heroicons/react/solid';
import { SkipIcon } from '@primer/octicons-react';
import {
  Accordion,
  DataTable,
  Empty,
  Flex,
  Grid,
  GridItem,
  Popover,
  Section,
  Tabs,
  Tag,
  Text,
  Tooltip,
} from '@redpanda-data/ui';
import { action, computed, makeObservable, observable } from 'mobx';
import {
  MdCheckCircleOutline,
  MdHourglassBottom,
  MdHourglassEmpty,
  MdLocalFireDepartment,
  MdOutlineQuiz,
  MdOutlineWarningAmber,
} from 'react-icons/md';
import { appGlobal } from '../../../state/appGlobal';
import { api } from '../../../state/backendApi';
import type { GroupDescription, GroupMemberDescription } from '../../../state/restInterfaces';
import { Features } from '../../../state/supportedFeatures';
import { uiSettings } from '../../../state/ui';
import { Button, DefaultSkeleton, IconButton, OptionGroup, numberToThousandsString } from '../../../utils/tsxUtils';
import PageContent from '../../misc/PageContent';
import { ShortNum } from '../../misc/ShortNum';
import { Statistic } from '../../misc/Statistic';
import { PageComponent, type PageInitHelper } from '../Page';
import AclList from '../topics/Tab.Acl/AclList';
import { DeleteOffsetsModal, EditOffsetsModal, type GroupOffset } from './Modals';
import type { GroupDeletingMode } from './Modals';

@observer
class GroupDetails extends PageComponent<{ groupId: string }> {
  @observable viewMode: 'topic' | 'member' = 'topic';
  @observable filterMode: 'showAll' | 'withLag' = 'showAll';

  @observable edittingOffsets: GroupOffset[] | null = null;

  @observable deletingMode: GroupDeletingMode = 'group';
  @observable deletingOffsets: GroupOffset[] | null = null;

  constructor(p: any) {
    super(p);
    makeObservable(this);
  }

  initPage(p: PageInitHelper): void {
    const group = decodeURIComponent(this.props.groupId);

    p.title = this.props.groupId;
    p.addBreadcrumb('Consumer Groups', '/groups');
    if (group)
      p.addBreadcrumb(group, `/${group}`, undefined, {
        canBeCopied: true,
        canBeTruncated: true,
      });

    this.refreshData(true);
    appGlobal.onRefresh = () => this.refreshData(true);
  }

  refreshData(force: boolean) {
    const group = decodeURIComponent(this.props.groupId);
    api.refreshConsumerGroup(group, force);
    api.refreshConsumerGroupAcls(group, force);
  }

  renderPartitions(group: GroupDescription) {
    return (
      <>
        <div style={{ display: 'flex', marginLeft: '.5em', marginBottom: '2em', gap: '1em', alignItems: 'flex-end' }}>
          <OptionGroup
            label="View"
            options={{
              Members: 'member',
              Topics: 'topic',
            }}
            value={this.viewMode}
            onChange={(s) => (this.viewMode = s)}
          />

          <OptionGroup
            label="Filter"
            options={{
              'Show All': 'showAll',
              'With Lag': 'withLag',
            }}
            value={this.filterMode}
            onChange={(s) => (this.filterMode = s)}
          />

          <span style={{ marginLeft: 'auto' }} />

          <Button variant="outline" onClick={() => this.editGroup()} disabledReason={cannotEditGroupReason(group)}>
            Edit Group
          </Button>
          <Button
            variant="outline"
            colorScheme="red"
            onClick={() => this.deleteGroup()}
            disabledReason={cannotDeleteGroupReason(group)}
          >
            Delete Group
          </Button>
        </div>

        {/* Main Content */}
        {this.viewMode === 'member' ? (
          <GroupByMembers group={group} onlyShowPartitionsWithLag={this.filterMode === 'withLag'} />
        ) : (
          <GroupByTopics
            group={group}
            onlyShowPartitionsWithLag={this.filterMode === 'withLag'}
            onEditOffsets={(g) => (this.edittingOffsets = g)}
            onDeleteOffsets={(offsets, mode) => {
              this.deletingMode = mode;
              this.deletingOffsets = offsets;
            }}
          />
        )}
      </>
    );
  }

  render() {
    // Get info about the group
    if (api.consumerGroups.size === 0) return DefaultSkeleton;
    const group = this.group;
    if (!group) return DefaultSkeleton;

    // Get info about each topic
    const totalPartitions = group.members.flatMap((m) => m.assignments).sum((a) => a.partitionIds.length);

    return (
      <PageContent className="groupDetails">
        {/* Statistics Card */}
        {uiSettings.consumerGroupDetails.showStatisticsBar && (
          <Section py={4}>
            <div className="statisticsBar">
              <Flex gap="2rem">
                <Statistic title="State" value={<GroupState group={group} />} />
                <Statistic title="Assigned Partitions" value={totalPartitions} />
                <ProtocolType group={group} />
                <Statistic title="Protocol Type" value={group.protocolType} />
                <Statistic title="Coordinator ID" value={group.coordinatorId} />
                <Statistic title="Total Lag" value={group.lagSum} />
              </Flex>
            </div>
          </Section>
        )}

        {/* Main Card */}
        <Section>
          {/* View Buttons */}
          <Tabs
            isFitted
            items={[
              {
                key: 'partitions',
                name: 'Partitions',
                component: this.renderPartitions(group),
              },
              {
                key: 'acl',
                name: 'ACL',
                component: <AclList acl={api.consumerGroupAcls.get(group.groupId)} />,
              },
            ]}
          />
        </Section>

        {/* Modals */}
        <>
          <EditOffsetsModal
            group={group}
            offsets={this.edittingOffsets}
            onClose={() => (this.edittingOffsets = null)}
          />

          <DeleteOffsetsModal
            group={group}
            mode={this.deletingMode}
            offsets={this.deletingOffsets}
            onClose={() => (this.deletingOffsets = null)}
          />
        </>
      </PageContent>
    );
  }

  @computed get group() {
    const groupId = decodeURIComponent(this.props.groupId);
    return api.consumerGroups.get(groupId);
  }

  @action editGroup() {
    const groupOffsets = this.group?.topicOffsets.flatMap((x) => {
      return x.partitionOffsets.map((p) => {
        return { topicName: x.topic, partitionId: p.partitionId, offset: p.groupOffset } as GroupOffset;
      });
    });

    if (!groupOffsets) return;

    this.edittingOffsets = groupOffsets;
  }

  @action deleteGroup() {
    const groupOffsets = this.group?.topicOffsets.flatMap((x) => {
      return x.partitionOffsets.map((p) => {
        return { topicName: x.topic, partitionId: p.partitionId, offset: p.groupOffset } as GroupOffset;
      });
    });

    if (!groupOffsets) return;

    this.deletingOffsets = groupOffsets;
    this.deletingMode = 'group';
  }
}

@observer
class GroupByTopics extends Component<{
  group: GroupDescription;
  onlyShowPartitionsWithLag: boolean;
  onEditOffsets: (offsets: GroupOffset[]) => void;
  onDeleteOffsets: (offsets: GroupOffset[], mode: GroupDeletingMode) => void;
}> {
  render() {
    const topicLags = this.props.group.topicOffsets;
    const p = this.props;
    const allAssignments = p.group.members.flatMap((m) =>
      m.assignments.map((as) => ({ member: m, topicName: as.topicName, partitions: as.partitionIds })),
    );

    const lagsFlat = topicLags.flatMap((topicLag) =>
      topicLag.partitionOffsets.map((partLag) => {
        const assignedMember = allAssignments.find(
          (e) => e.topicName === topicLag.topic && e.partitions.includes(partLag.partitionId),
        );

        return {
          topicName: topicLag.topic,
          partitionId: partLag.partitionId,
          groupOffset: partLag.groupOffset,
          highWaterMark: partLag.highWaterMark,
          lag: partLag.lag,

          assignedMember: assignedMember?.member,
          id: assignedMember?.member.id,
          clientId: assignedMember?.member.clientId,
          host: assignedMember?.member.clientHost,
        };
      }),
    );

    const lagGroupsByTopic = lagsFlat
      .groupInto((e) => e.topicName)
      .sort((a, b) => a.key.localeCompare(b.key))
      .map((x) => ({ topicName: x.key, partitions: x.items }));

    const topicEntries = lagGroupsByTopic.map((g) => {
      const totalLagAll = g.partitions.sum((c) => c.lag ?? 0);
      const partitionsAssigned = g.partitions.filter((c) => c.assignedMember).length;

      if (p.onlyShowPartitionsWithLag) g.partitions.removeAll((e) => e.lag === 0);

      if (g.partitions.length === 0) return null;

      return {
        heading: (
          <Flex justifyContent="space-between">
            <Flex alignItems="center" gap={1}>
              {/* Title */}
              <Text fontWeight={600} fontSize="lg">
                {g.topicName}
              </Text>

              {/* EditButtons */}
              <div style={{ width: '2px' }} />

              <IconButton
                onClick={(e) => {
                  p.onEditOffsets(g.partitions);
                  e.stopPropagation();
                }}
                disabledReason={cannotEditGroupReason(this.props.group)}
              >
                <PencilIcon />
              </IconButton>
              <IconButton
                onClick={(e) => {
                  p.onDeleteOffsets(g.partitions, 'topic');
                  e.stopPropagation();
                }}
                disabledReason={cannotDeleteGroupOffsetsReason(this.props.group)}
              >
                <TrashIcon />
              </IconButton>

              <Flex ml={2} gap={1}>
                {/* InfoTags */}
                <Tooltip placement="top" label="Summed lag of all partitions of the topic" hasArrow>
                  <Tag colorScheme="brand" variant="solid">
                    lag: {numberToThousandsString(totalLagAll)}
                  </Tag>
                </Tooltip>
                <Tooltip placement="top" label="Number of assigned partitions" hasArrow>
                  <Tag colorScheme="brand" variant="solid">
                    assigned partitions: {partitionsAssigned}
                  </Tag>
                </Tooltip>
              </Flex>
            </Flex>
            <Button
              variant="outline"
              size="sm"
              onClick={() => appGlobal.history.push(`/topics/${encodeURIComponent(g.topicName)}`)}
            >
              View Topic
            </Button>
          </Flex>
        ),
        description: (
          <DataTable<{
            topicName: string;
            partitionId: number;
            groupOffset: number;
            highWaterMark: number;
            lag: number;
            assignedMember: GroupMemberDescription | undefined;
            id: string | undefined;
            clientId: string | undefined;
            host: string | undefined;
          }>
            pagination
            sorting
            data={g.partitions}
            columns={[
              {
                size: 100,
                header: 'Partition',
                accessorKey: 'partitionId',
              },
              {
                size: Number.POSITIVE_INFINITY,
                header: 'Assigned Member',
                accessorKey: 'id',
                cell: ({
                  row: {
                    original: { assignedMember, id, clientId },
                  },
                }) =>
                  assignedMember ? (
                    renderMergedID(id, clientId)
                  ) : (
                    <span style={{ margin: '0 3px' }}>
                      <SkipIcon /> No assigned member
                    </span>
                  ),
              },
              {
                header: 'Host',
                accessorKey: 'host',
                cell: ({
                  row: {
                    original: { host },
                  },
                }) =>
                  host ?? (
                    <span style={{ opacity: 0.66, margin: '0 3px' }}>
                      <SkipIcon />
                    </span>
                  ),
              },
              {
                size: 120,
                header: 'Log End Offset',
                accessorKey: 'highWaterMark',
                cell: ({ row: { original } }) => numberToThousandsString(original.highWaterMark),
              },
              {
                size: 120,
                header: 'Group Offset',
                accessorKey: 'groupOffset',
                cell: ({ row: { original } }) => numberToThousandsString(original.groupOffset),
              },
              {
                size: 80,
                header: 'Lag',
                accessorKey: 'lag',
                cell: ({ row: { original } }) => ShortNum({ value: original.lag, tooltip: true }),
              },
              {
                size: 1,
                header: '',
                id: 'action',
                cell: ({ row: { original } }) => (
                  <Flex pr={2} gap={1}>
                    <IconButton
                      onClick={() => p.onEditOffsets([original])}
                      disabledReason={cannotEditGroupReason(this.props.group)}
                    >
                      <PencilIcon />
                    </IconButton>
                    <IconButton
                      onClick={() => p.onDeleteOffsets([original], 'partition')}
                      disabledReason={cannotDeleteGroupOffsetsReason(this.props.group)}
                    >
                      <TrashIcon />
                    </IconButton>
                  </Flex>
                ),
              },
            ]}
          />
        ),
      };
    });

    const defaultExpand: number | undefined =
      lagGroupsByTopic.length === 1
        ? 0 // only one -> expand
        : undefined; // more than one -> collapse

    const nullEntries = topicEntries.filter((e) => e == null).length;
    if (topicEntries.length === 0 || topicEntries.length === nullEntries) {
      return (
        <Empty
          description={
            p.onlyShowPartitionsWithLag ? (
              <span>All {topicEntries.length} topics have been filtered (no lag on any partition).</span>
            ) : (
              'No data found'
            )
          }
        />
      );
    }

    return <Accordion items={topicEntries.filterNull()} defaultIndex={defaultExpand} />;
  }
}

@observer
class GroupByMembers extends Component<{ group: GroupDescription; onlyShowPartitionsWithLag: boolean }> {
  render() {
    const topicLags = this.props.group.topicOffsets;
    const p = this.props;

    const memberEntries = p.group.members
      // sorting actually not necessary
      // .sort((a, b) => a.id.localeCompare(b.id))
      .map((m) => {
        const assignments = m.assignments;

        const assignmentsFlat = assignments.flatMap((a) =>
          a.partitionIds.map((id) => {
            const topicLag = topicLags.find((t) => t.topic === a.topicName);
            const partLag = topicLag?.partitionOffsets.find((p) => p.partitionId === id)?.lag;
            return {
              topicName: a.topicName,
              partitionId: id,
              partitionLag: partLag ?? 0,
            };
          }),
        );

        const totalLag = assignmentsFlat.sum((t) => t.partitionLag ?? 0);
        const totalPartitions = assignmentsFlat.length;

        if (p.onlyShowPartitionsWithLag) assignmentsFlat.removeAll((e) => e.partitionLag === 0);

        if (assignmentsFlat.length === 0) return null;

        return {
          heading: (
            <Flex alignItems="baseline" gap="1">
              <Text fontWeight={600} fontSize="lg" mr={4}>
                {renderMergedID(m.id, m.clientId)}
              </Text>
              <Tooltip placement="top" label="Host of the member" hasArrow>
                <Tag colorScheme="blue">host: {m.clientHost}</Tag>
              </Tooltip>
              <Tooltip placement="top" label="Number of assigned partitions" hasArrow>
                <Tag variant="solid" colorScheme="brand">
                  partitions: {totalPartitions}
                </Tag>
              </Tooltip>
              <Tooltip placement="top" label="Summed lag over all assigned partitions of all topics" hasArrow>
                <Tag variant="solid" colorScheme="brand">
                  lag: {totalLag}
                </Tag>
              </Tooltip>
            </Flex>
          ),
          description: (
            <DataTable<{ topicName: string; partitionId: number; partitionLag: number }>
              pagination
              sorting
              data={assignmentsFlat}
              columns={[
                {
                  size: 130,
                  header: 'Topic',
                  accessorKey: 'topicName',
                  cell: ({
                    row: {
                      original: { topicName },
                    },
                  }) => (
                    <div
                      className="hoverLink"
                      onClick={() => appGlobal.history.push(`/topics/${encodeURIComponent(topicName)}`)}
                    >
                      {topicName}
                    </div>
                  ),
                },
                {
                  header: 'Partition',
                  accessorKey: 'partitionId',
                },
                {
                  header: 'Lag',
                  accessorKey: 'partitionLag',
                  cell: ({
                    row: {
                      original: { partitionLag },
                    },
                  }) => numberToThousandsString(partitionLag),
                },
              ]}
            />
          ),
        };
      });

    const defaultExpandIndex: number | undefined =
      p.group.members.length === 1
        ? 0 // if only one entry, expand it
        : undefined; // more than one -> collapse

    const nullEntries = memberEntries.filter((e) => e == null).length;
    if (memberEntries.length === 0 || memberEntries.length === nullEntries) {
      return (
        <Empty
          description={
            p.onlyShowPartitionsWithLag ? (
              <span>All {memberEntries.length} members have been filtered (no lag on any partition).</span>
            ) : (
              'No data found'
            )
          }
        />
      );
    }

    return <Accordion items={memberEntries.filterNull()} defaultIndex={defaultExpandIndex} />;
  }
}

const renderMergedID = (id?: string, clientId?: string) => {
  if (clientId && id?.startsWith(clientId)) {
    // should always be true...
    const suffix = id.substring(clientId.length);

    return (
      <span className="consumerGroupCompleteID">
        <span className="consumerGroupName">{clientId}</span>
        <span className="consumerGroupSuffix">{suffix}</span>
      </span>
    );
  }
  // A client might be connected but it hasn't any assignments yet because it just joined the group
  if (clientId) {
    return <span className="consumerGroupCompleteID">{clientId ?? id ?? ''}</span>;
  }

  return null;
};

type StateIcon = 'stable' | 'completingrebalance' | 'preparingrebalance' | 'empty' | 'dead' | 'unknown';

const stateIcons = new Map<StateIcon, JSX.Element>([
  ['stable', <MdCheckCircleOutline key="stable" size={16} color="#52c41a" />],
  ['completingrebalance', <MdHourglassBottom key="completingrebalance" size={16} color="#52c41a" />],
  ['preparingrebalance', <MdHourglassEmpty key="preparingrebalance" size={16} color="orange" />],
  ['empty', <MdOutlineWarningAmber key="empty" size={16} color="orange" />],
  ['dead', <MdLocalFireDepartment key="dead" size={16} color="orangered" />],
  ['unknown', <MdOutlineQuiz key="unknown" size={16} />],
]);

const stateIconNames: Record<StateIcon, string> = {
  stable: 'Stable',
  completingrebalance: 'Completing Rebalance',
  preparingrebalance: 'Preparing Rebalance',
  empty: 'Empty',
  dead: 'Dead',
  unknown: 'Unknown',
};

const stateIconDescriptions: Record<StateIcon, string> = {
  stable: 'Consumer group has members which have been assigned partitions',
  completingrebalance: 'Kafka is assigning partitions to group members',
  preparingrebalance: 'A reassignment of partitions is required, members have been asked to stop consuming',
  empty: 'Consumer group exists, but does not have any members',
  dead: 'Consumer group does not have any members and its metadata has been removed',
  unknown: 'Group state is not known',
};

const consumerGroupStateTable = (
  <Grid templateColumns="auto 300px" gap={4}>
    {Array.from(stateIcons.entries()).map(([key, icon]) => (
      <React.Fragment key={key}>
        {/* Icon column */}
        <GridItem display="flex" alignItems="center" gap={2}>
          {icon} <strong>{stateIconNames[key]}</strong>
        </GridItem>

        {/* Description column */}
        <GridItem>{stateIconDescriptions[key]}</GridItem>
      </React.Fragment>
    ))}
  </Grid>
);

export const GroupState = (p: { group: GroupDescription }) => {
  const state = p.group.state.toLowerCase();
  const icon = stateIcons.get(state as StateIcon);

  return (
    <Popover isInPortal trigger="hover" size="auto" placement="right" hideCloseButton content={consumerGroupStateTable}>
      <Flex gap={2} alignItems="center">
        {icon}
        <span> {p.group.state}</span>
      </Flex>
    </Popover>
  );
};
const ProtocolType = (p: { group: GroupDescription }) => {
  const protocol = p.group.protocolType;
  if (protocol === 'consumer') return null;

  return <Statistic title="Protocol" value={protocol} />;
};

function cannotEditGroupReason(group: GroupDescription): string | undefined {
  if (group.noEditPerms) return "You don't have 'editConsumerGroup' permissions for this group";
  if (group.isInUse) return 'Consumer groups with active members cannot be edited';
  if (!Features.patchGroup) return 'This cluster does not support editting group offsets';
}

function cannotDeleteGroupReason(group: GroupDescription): string | undefined {
  if (group.noDeletePerms) return "You don't have 'deleteConsumerGroup' permissions for this group";
  if (group.isInUse) return 'Consumer groups with active members cannot be deleted';
  if (!Features.deleteGroup) return 'This cluster does not support deleting groups';
}

function cannotDeleteGroupOffsetsReason(group: GroupDescription): string | undefined {
  if (group.noEditPerms) return "You don't have 'deleteConsumerGroup' permissions for this group";
  if (group.isInUse) return 'Consumer groups with active members cannot be deleted';
  if (!Features.deleteGroupOffsets) return 'This cluster does not support deleting group offsets';
}

export default GroupDetails;
